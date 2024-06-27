import * as Yup from 'yup';
import humanizeString from 'humanize-string';

declare module 'yup' {
  // eslint-disable-next-line @typescript-eslint/ban-types
  interface MixedSchema<TType extends Yup.Maybe<{}>> {
    email(): MixedSchema<TType>;
    max(minValue: number): MixedSchema<TType>;
    min(maxValue: number): MixedSchema<TType>;
    required(): MixedSchema<TType>;
    url(): MixedSchema<TType>;
  }

  interface ArraySchema<
    TIn extends any[] | null | undefined,
    TContext,
    TDefault = undefined,
    TFlags extends Yup.Flags = '',
  > extends Yup.Schema<TIn, TContext, TDefault, TFlags> {
    unique(
      message?: string,
      mapper?: (a: any) => any,
    ): ArraySchema<TIn, TContext, TDefault, TFlags>;
  }
}

const emailRegex =
  /^[\w!#$%&'*+./=?^`{|}~-]+@[\dA-Za-z](?:[\dA-Za-z-]{0,61}[\dA-Za-z])?(?:\.[\dA-Za-z](?:[\dA-Za-z-]{0,61}[\dA-Za-z])?)*$/;

Yup.addMethod(Yup.string, 'email', function (errorMessage) {
  return this.test(`email`, errorMessage, function (value) {
    if (!value) {
      return true;
    }

    return (
      emailRegex.test(value || '') || this.createError({ message: errorMessage, path: this.path })
    );
  });
});

Yup.addMethod(Yup.array, 'unique', function (message, mapper = (a: any) => a) {
  return this.test('unique', message, function (list = []) {
    return list.length === new Set(list.map((element) => mapper(element))).size;
  });
});

export function convertStringToSlug(value: string) {
  let slug = value
    .toLowerCase()
    .replaceAll(/[^\da-z-]+/g, '-')
    .replaceAll(/-{2,}/g, '-')
    .replaceAll(/^-+|-+$/g, '');

  if (!/^[a-z]/.test(slug)) {
    slug = `a${slug}`;
  }

  return slug;
}

export function getYupValidationResolver(validationSchema: Yup.AnyObjectSchema) {
  return async (data: any) => {
    try {
      const values = await validationSchema.validate(data, {
        abortEarly: false,
      });

      return {
        errors: {},
        values,
      };
    } catch (error: unknown) {
      const err = error as Yup.ValidationError;

      return {
        errors: Object.fromEntries(
          err.inner.map((currentError) => {
            return [
              currentError.path,
              {
                message: getFieldErrorMessage(currentError),
                type: currentError.type ?? 'validation',
              },
            ];
          }),
        ),
        values: {},
      };
    }
  };
}

Yup.setLocale({
  mixed: {
    notType: '',
    required: '',
  },
  number: {
    integer: '',
    max: '',
    min: '',
  },
  string: {
    email: '',
    max: '',
    min: '',
  },
});

function capitalizeText(string: string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}

function getFieldErrorMessage(error: Yup.ValidationError): string {
  let fieldLabel =
    (error.params?.label as string) ||
    humanizeString(error.path?.split('.').pop() || '') ||
    'value';
  fieldLabel = capitalizeText(fieldLabel);
  let errorMessage = (error.message || '').replaceAll('{{fieldLabel}}', fieldLabel);
  errorMessage = capitalizeText(errorMessage);

  if (error.type === 'required' || error.type === 'optionality') {
    return errorMessage || `${fieldLabel} cannot be empty`;
  }

  if (error.type === 'email') {
    return `${fieldLabel} must be a valid email address`;
  }

  if (error.type === 'typeError' && error.params?.type === 'number') {
    return errorMessage || `${fieldLabel} must be a valid number`;
  }

  if (error.type === 'integer') {
    return errorMessage || `${fieldLabel} must be a valid integer`;
  }

  if (!error.params) {
    return errorMessage;
  }

  if (
    error.type === 'min' &&
    typeof error.params.value === 'number' &&
    typeof error.params.min === 'number'
  ) {
    return errorMessage || `${fieldLabel} must be more than or equal to ${error.params.min}`;
  }

  if (
    error.type === 'max' &&
    typeof error.params.value === 'number' &&
    typeof error.params.max === 'number'
  ) {
    return errorMessage || `${fieldLabel} must be less than or equal to ${error.params.max}`;
  }

  if (
    error.type === 'max' &&
    typeof error.params.value === 'string' &&
    typeof error.params.max === 'number'
  ) {
    return errorMessage || `${fieldLabel} must contain at most ${error.params.max} characters`;
  }

  if (
    error.type === 'min' &&
    typeof error.params.value === 'string' &&
    typeof error.params.min === 'number'
  ) {
    return errorMessage || `${fieldLabel} must contain at least ${error.params.min} characters`;
  }

  return errorMessage;
}

export type YupInferType<T extends Yup.ISchema<any, any>> = Yup.InferType<T>;

// eslint-disable-next-line unicorn/prefer-export-from
export const yup = Yup;
