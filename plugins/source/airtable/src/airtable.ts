// Based on https://airtable.com/api/meta
export type APIBase = {
  id: string;
  name: string;
};

export enum APIFieldType {
  autoNumber = 'autoNumber',
  barcode = 'barcode',
  button = 'button',
  checkbox = 'checkbox',
  count = 'count',
  createdBy = 'createdBy',
  createdTime = 'createdTime',
  currency = 'currency',
  date = 'date',
  dateTime = 'dateTime',
  duration = 'duration',
  email = 'email',
  externalSyncSource = 'externalSyncSource',
  formula = 'formula',
  lastModifiedBy = 'lastModifiedBy',
  lastModifiedTime = 'lastModifiedTime',
  multilineText = 'multilineText',
  multipleAttachments = 'multipleAttachments',
  multipleCollaborators = 'multipleCollaborators',
  multipleLookupValues = 'multipleLookupValues',
  multipleRecordLinks = 'multipleRecordLinks',
  multipleSelects = 'multipleSelects',
  number = 'number',
  percent = 'percent',
  phoneNumber = 'phoneNumber',
  rating = 'rating',
  richText = 'richText',
  rollup = 'rollup',
  singleCollaborator = 'singleCollaborator',
  singleLineText = 'singleLineText',
  singleSelect = 'singleSelect',
  url = 'url',
}

export type APIBaseField = {
  id: string;
  name: string;
  description: string;
  type: APIFieldType;
};

export type APIFieldNoOptions = APIBaseField & {
  type:
    | APIFieldType.autoNumber
    | APIFieldType.barcode
    | APIFieldType.button
    | APIFieldType.createdBy
    | APIFieldType.email
    | APIFieldType.lastModifiedBy
    | APIFieldType.multilineText
    | APIFieldType.phoneNumber
    | APIFieldType.richText
    | APIFieldType.singleCollaborator
    | APIFieldType.singleLineText
    | APIFieldType.url;
  options: undefined;
};

export type APIFieldCheckbox = APIBaseField & {
  type: APIFieldType.checkbox;
  options: {
    icon: 'check' | 'xCheckbox' | 'star' | 'heart' | 'thumbsUp' | 'flag' | 'dot';
    color:
      | 'greenBright'
      | 'tealBright'
      | 'cyanBright'
      | 'blueBright'
      | 'purpleBright'
      | 'pinkBright'
      | 'redBright'
      | 'orangeBright'
      | 'yellowBright'
      | 'grayBright';
  };
};

export type APIFieldCount = APIBaseField & {
  type: APIFieldType.count;
  options: {
    // false when recordLinkFieldId is null, e.g. the referenced column was deleted.
    isValid: boolean;
    recordLinkFieldId: string | null;
  };
};

export type APIFieldCreatedTime = APIBaseField & {
  type: APIFieldType.createdTime;
  options: {
    result:
      | {
          type: APIFieldType.date;
          options: APIDateOptions;
        }
      | {
          type: APIFieldType.dateTime;
          options: APIDateTimeOptions;
        };
  };
};

export type APIFieldCurrency = APIBaseField & {
  type: APIFieldType.currency;
  options: {
    // from 0 to 7 inclusive
    precision: number;
    symbol: string;
  };
};

export type APIDateOptions = {
  dateFormat: {
    name: 'local' | 'friendly' | 'us' | 'european' | 'iso';
    format: 'l' | 'LL' | 'M/D/YYYY' | 'D/M/YYYY' | 'YYYY-MM-DD';
  };
};

export type APIDate = APIBaseField & {
  type: APIFieldType.date;
  options: APIDateOptions;
};

export type APIDateTimeOptions = {
  dateFormat: {
    name: 'local' | 'friendly' | 'us' | 'european' | 'iso';
    format: 'l' | 'LL' | 'M/D/YYYY' | 'D/M/YYYY' | 'YYYY-MM-DD';
  };
  timeFormat: {
    name: '12hour' | '24hour';
    format: 'h:mma' | 'HH:mm';
  };
  timeZone:
    | 'utc'
    | 'client'
    | 'Africa/Abidjan'
    | 'Africa/Accra'
    | 'Africa/Addis_Ababa'
    | 'Africa/Algiers'
    | 'Africa/Asmara'
    | 'Africa/Bamako'
    | 'Africa/Bangui'
    | 'Africa/Banjul'
    | 'Africa/Bissau'
    | 'Africa/Blantyre'
    | 'Africa/Brazzaville'
    | 'Africa/Bujumbura'
    | 'Africa/Cairo'
    | 'Africa/Casablanca'
    | 'Africa/Ceuta'
    | 'Africa/Conakry'
    | 'Africa/Dakar'
    | 'Africa/Dar_es_Salaam'
    | 'Africa/Djibouti'
    | 'Africa/Douala'
    | 'Africa/El_Aaiun'
    | 'Africa/Freetown'
    | 'Africa/Gaborone'
    | 'Africa/Harare'
    | 'Africa/Johannesburg'
    | 'Africa/Juba'
    | 'Africa/Kampala'
    | 'Africa/Khartoum'
    | 'Africa/Kigali'
    | 'Africa/Kinshasa'
    | 'Africa/Lagos'
    | 'Africa/Libreville'
    | 'Africa/Lome'
    | 'Africa/Luanda'
    | 'Africa/Lubumbashi'
    | 'Africa/Lusaka'
    | 'Africa/Malabo'
    | 'Africa/Maputo'
    | 'Africa/Maseru'
    | 'Africa/Mbabane'
    | 'Africa/Mogadishu'
    | 'Africa/Monrovia'
    | 'Africa/Nairobi'
    | 'Africa/Ndjamena'
    | 'Africa/Niamey'
    | 'Africa/Nouakchott'
    | 'Africa/Ouagadougou'
    | 'Africa/Porto-Novo'
    | 'Africa/Sao_Tome'
    | 'Africa/Tripoli'
    | 'Africa/Tunis'
    | 'Africa/Windhoek'
    | 'America/Adak'
    | 'America/Anchorage'
    | 'America/Anguilla'
    | 'America/Antigua'
    | 'America/Araguaina'
    | 'America/Argentina/Buenos_Aires'
    | 'America/Argentina/Catamarca'
    | 'America/Argentina/Cordoba'
    | 'America/Argentina/Jujuy'
    | 'America/Argentina/La_Rioja'
    | 'America/Argentina/Mendoza'
    | 'America/Argentina/Rio_Gallegos'
    | 'America/Argentina/Salta'
    | 'America/Argentina/San_Juan'
    | 'America/Argentina/San_Luis'
    | 'America/Argentina/Tucuman'
    | 'America/Argentina/Ushuaia'
    | 'America/Aruba'
    | 'America/Asuncion'
    | 'America/Atikokan'
    | 'America/Bahia'
    | 'America/Bahia_Banderas'
    | 'America/Barbados'
    | 'America/Belem'
    | 'America/Belize'
    | 'America/Blanc-Sablon'
    | 'America/Boa_Vista'
    | 'America/Bogota'
    | 'America/Boise'
    | 'America/Cambridge_Bay'
    | 'America/Campo_Grande'
    | 'America/Cancun'
    | 'America/Caracas'
    | 'America/Cayenne'
    | 'America/Cayman'
    | 'America/Chicago'
    | 'America/Chihuahua'
    | 'America/Costa_Rica'
    | 'America/Creston'
    | 'America/Cuiaba'
    | 'America/Curacao'
    | 'America/Danmarkshavn'
    | 'America/Dawson'
    | 'America/Dawson_Creek'
    | 'America/Denver'
    | 'America/Detroit'
    | 'America/Dominica'
    | 'America/Edmonton'
    | 'America/Eirunepe'
    | 'America/El_Salvador'
    | 'America/Fort_Nelson'
    | 'America/Fortaleza'
    | 'America/Glace_Bay'
    | 'America/Godthab'
    | 'America/Goose_Bay'
    | 'America/Grand_Turk'
    | 'America/Grenada'
    | 'America/Guadeloupe'
    | 'America/Guatemala'
    | 'America/Guayaquil'
    | 'America/Guyana'
    | 'America/Halifax'
    | 'America/Havana'
    | 'America/Hermosillo'
    | 'America/Indiana/Indianapolis'
    | 'America/Indiana/Knox'
    | 'America/Indiana/Marengo'
    | 'America/Indiana/Petersburg'
    | 'America/Indiana/Tell_City'
    | 'America/Indiana/Vevay'
    | 'America/Indiana/Vincennes'
    | 'America/Indiana/Winamac'
    | 'America/Inuvik'
    | 'America/Iqaluit'
    | 'America/Jamaica'
    | 'America/Juneau'
    | 'America/Kentucky/Louisville'
    | 'America/Kentucky/Monticello'
    | 'America/Kralendijk'
    | 'America/La_Paz'
    | 'America/Lima'
    | 'America/Los_Angeles'
    | 'America/Lower_Princes'
    | 'America/Maceio'
    | 'America/Managua'
    | 'America/Manaus'
    | 'America/Marigot'
    | 'America/Martinique'
    | 'America/Matamoros'
    | 'America/Mazatlan'
    | 'America/Menominee'
    | 'America/Merida'
    | 'America/Metlakatla'
    | 'America/Mexico_City'
    | 'America/Miquelon'
    | 'America/Moncton'
    | 'America/Monterrey'
    | 'America/Montevideo'
    | 'America/Montserrat'
    | 'America/Nassau'
    | 'America/New_York'
    | 'America/Nipigon'
    | 'America/Nome'
    | 'America/Noronha'
    | 'America/North_Dakota/Beulah'
    | 'America/North_Dakota/Center'
    | 'America/North_Dakota/New_Salem'
    | 'America/Nuuk'
    | 'America/Ojinaga'
    | 'America/Panama'
    | 'America/Pangnirtung'
    | 'America/Paramaribo'
    | 'America/Phoenix'
    | 'America/Port-au-Prince'
    | 'America/Port_of_Spain'
    | 'America/Porto_Velho'
    | 'America/Puerto_Rico'
    | 'America/Punta_Arenas'
    | 'America/Rainy_River'
    | 'America/Rankin_Inlet'
    | 'America/Recife'
    | 'America/Regina'
    | 'America/Resolute'
    | 'America/Rio_Branco'
    | 'America/Santarem'
    | 'America/Santiago'
    | 'America/Santo_Domingo'
    | 'America/Sao_Paulo'
    | 'America/Scoresbysund'
    | 'America/Sitka'
    | 'America/St_Barthelemy'
    | 'America/St_Johns'
    | 'America/St_Kitts'
    | 'America/St_Lucia'
    | 'America/St_Thomas'
    | 'America/St_Vincent'
    | 'America/Swift_Current'
    | 'America/Tegucigalpa'
    | 'America/Thule'
    | 'America/Thunder_Bay'
    | 'America/Tijuana'
    | 'America/Toronto'
    | 'America/Tortola'
    | 'America/Vancouver'
    | 'America/Whitehorse'
    | 'America/Winnipeg'
    | 'America/Yakutat'
    | 'America/Yellowknife'
    | 'Antarctica/Casey'
    | 'Antarctica/Davis'
    | 'Antarctica/DumontDUrville'
    | 'Antarctica/Macquarie'
    | 'Antarctica/Mawson'
    | 'Antarctica/McMurdo'
    | 'Antarctica/Palmer'
    | 'Antarctica/Rothera'
    | 'Antarctica/Syowa'
    | 'Antarctica/Troll'
    | 'Antarctica/Vostok'
    | 'Arctic/Longyearbyen'
    | 'Asia/Aden'
    | 'Asia/Almaty'
    | 'Asia/Amman'
    | 'Asia/Anadyr'
    | 'Asia/Aqtau'
    | 'Asia/Aqtobe'
    | 'Asia/Ashgabat'
    | 'Asia/Atyrau'
    | 'Asia/Baghdad'
    | 'Asia/Bahrain'
    | 'Asia/Baku'
    | 'Asia/Bangkok'
    | 'Asia/Barnaul'
    | 'Asia/Beirut'
    | 'Asia/Bishkek'
    | 'Asia/Brunei'
    | 'Asia/Chita'
    | 'Asia/Choibalsan'
    | 'Asia/Colombo'
    | 'Asia/Damascus'
    | 'Asia/Dhaka'
    | 'Asia/Dili'
    | 'Asia/Dubai'
    | 'Asia/Dushanbe'
    | 'Asia/Famagusta'
    | 'Asia/Gaza'
    | 'Asia/Hebron'
    | 'Asia/Ho_Chi_Minh'
    | 'Asia/Hong_Kong'
    | 'Asia/Hovd'
    | 'Asia/Irkutsk'
    | 'Asia/Istanbul'
    | 'Asia/Jakarta'
    | 'Asia/Jayapura'
    | 'Asia/Jerusalem'
    | 'Asia/Kabul'
    | 'Asia/Kamchatka'
    | 'Asia/Karachi'
    | 'Asia/Kathmandu'
    | 'Asia/Khandyga'
    | 'Asia/Kolkata'
    | 'Asia/Krasnoyarsk'
    | 'Asia/Kuala_Lumpur'
    | 'Asia/Kuching'
    | 'Asia/Kuwait'
    | 'Asia/Macau'
    | 'Asia/Magadan'
    | 'Asia/Makassar'
    | 'Asia/Manila'
    | 'Asia/Muscat'
    | 'Asia/Nicosia'
    | 'Asia/Novokuznetsk'
    | 'Asia/Novosibirsk'
    | 'Asia/Omsk'
    | 'Asia/Oral'
    | 'Asia/Phnom_Penh'
    | 'Asia/Pontianak'
    | 'Asia/Pyongyang'
    | 'Asia/Qatar'
    | 'Asia/Qostanay'
    | 'Asia/Qyzylorda'
    | 'Asia/Rangoon'
    | 'Asia/Riyadh'
    | 'Asia/Sakhalin'
    | 'Asia/Samarkand'
    | 'Asia/Seoul'
    | 'Asia/Shanghai'
    | 'Asia/Singapore'
    | 'Asia/Srednekolymsk'
    | 'Asia/Taipei'
    | 'Asia/Tashkent'
    | 'Asia/Tbilisi'
    | 'Asia/Tehran'
    | 'Asia/Thimphu'
    | 'Asia/Tokyo'
    | 'Asia/Tomsk'
    | 'Asia/Ulaanbaatar'
    | 'Asia/Urumqi'
    | 'Asia/Ust-Nera'
    | 'Asia/Vientiane'
    | 'Asia/Vladivostok'
    | 'Asia/Yakutsk'
    | 'Asia/Yangon'
    | 'Asia/Yekaterinburg'
    | 'Asia/Yerevan'
    | 'Atlantic/Azores'
    | 'Atlantic/Bermuda'
    | 'Atlantic/Canary'
    | 'Atlantic/Cape_Verde'
    | 'Atlantic/Faroe'
    | 'Atlantic/Madeira'
    | 'Atlantic/Reykjavik'
    | 'Atlantic/South_Georgia'
    | 'Atlantic/St_Helena'
    | 'Atlantic/Stanley'
    | 'Australia/Adelaide'
    | 'Australia/Brisbane'
    | 'Australia/Broken_Hill'
    | 'Australia/Currie'
    | 'Australia/Darwin'
    | 'Australia/Eucla'
    | 'Australia/Hobart'
    | 'Australia/Lindeman'
    | 'Australia/Lord_Howe'
    | 'Australia/Melbourne'
    | 'Australia/Perth'
    | 'Australia/Sydney'
    | 'Europe/Amsterdam'
    | 'Europe/Andorra'
    | 'Europe/Astrakhan'
    | 'Europe/Athens'
    | 'Europe/Belgrade'
    | 'Europe/Berlin'
    | 'Europe/Bratislava'
    | 'Europe/Brussels'
    | 'Europe/Bucharest'
    | 'Europe/Budapest'
    | 'Europe/Busingen'
    | 'Europe/Chisinau'
    | 'Europe/Copenhagen'
    | 'Europe/Dublin'
    | 'Europe/Gibraltar'
    | 'Europe/Guernsey'
    | 'Europe/Helsinki'
    | 'Europe/Isle_of_Man'
    | 'Europe/Istanbul'
    | 'Europe/Jersey'
    | 'Europe/Kaliningrad'
    | 'Europe/Kiev'
    | 'Europe/Kirov'
    | 'Europe/Lisbon'
    | 'Europe/Ljubljana'
    | 'Europe/London'
    | 'Europe/Luxembourg'
    | 'Europe/Madrid'
    | 'Europe/Malta'
    | 'Europe/Mariehamn'
    | 'Europe/Minsk'
    | 'Europe/Monaco'
    | 'Europe/Moscow'
    | 'Europe/Nicosia'
    | 'Europe/Oslo'
    | 'Europe/Paris'
    | 'Europe/Podgorica'
    | 'Europe/Prague'
    | 'Europe/Riga'
    | 'Europe/Rome'
    | 'Europe/Samara'
    | 'Europe/San_Marino'
    | 'Europe/Sarajevo'
    | 'Europe/Saratov'
    | 'Europe/Simferopol'
    | 'Europe/Skopje'
    | 'Europe/Sofia'
    | 'Europe/Stockholm'
    | 'Europe/Tallinn'
    | 'Europe/Tirane'
    | 'Europe/Ulyanovsk'
    | 'Europe/Uzhgorod'
    | 'Europe/Vaduz'
    | 'Europe/Vatican'
    | 'Europe/Vienna'
    | 'Europe/Vilnius'
    | 'Europe/Volgograd'
    | 'Europe/Warsaw'
    | 'Europe/Zagreb'
    | 'Europe/Zaporozhye'
    | 'Europe/Zurich'
    | 'Indian/Antananarivo'
    | 'Indian/Chagos'
    | 'Indian/Christmas'
    | 'Indian/Cocos'
    | 'Indian/Comoro'
    | 'Indian/Kerguelen'
    | 'Indian/Mahe'
    | 'Indian/Maldives'
    | 'Indian/Mauritius'
    | 'Indian/Mayotte'
    | 'Indian/Reunion'
    | 'Pacific/Apia'
    | 'Pacific/Auckland'
    | 'Pacific/Bougainville'
    | 'Pacific/Chatham'
    | 'Pacific/Chuuk'
    | 'Pacific/Easter'
    | 'Pacific/Efate'
    | 'Pacific/Enderbury'
    | 'Pacific/Fakaofo'
    | 'Pacific/Fiji'
    | 'Pacific/Funafuti'
    | 'Pacific/Galapagos'
    | 'Pacific/Gambier'
    | 'Pacific/Guadalcanal'
    | 'Pacific/Guam'
    | 'Pacific/Honolulu'
    | 'Pacific/Kanton'
    | 'Pacific/Kiritimati'
    | 'Pacific/Kosrae'
    | 'Pacific/Kwajalein'
    | 'Pacific/Majuro'
    | 'Pacific/Marquesas'
    | 'Pacific/Midway'
    | 'Pacific/Nauru'
    | 'Pacific/Niue'
    | 'Pacific/Norfolk'
    | 'Pacific/Noumea'
    | 'Pacific/Pago_Pago'
    | 'Pacific/Palau'
    | 'Pacific/Pitcairn'
    | 'Pacific/Pohnpei'
    | 'Pacific/Port_Moresby'
    | 'Pacific/Rarotonga'
    | 'Pacific/Saipan'
    | 'Pacific/Tahiti'
    | 'Pacific/Tarawa'
    | 'Pacific/Tongatapu'
    | 'Pacific/Wake'
    | 'Pacific/Wallis';
};

export type APIDateTime = APIBaseField & {
  type: APIFieldType.dateTime;
  options: APIDateTimeOptions;
};

export type APIDuration = APIBaseField & {
  type: APIFieldType.duration;
  options: {
    durationFormat: 'h:mm' | 'h:mm:ss' | 'h:mm:ss.S' | 'h:mm:ss.SS' | 'h:mm:ss.SSS';
  };
};

export type APIExternalSyncSource = APIBaseField & {
  type: APIFieldType.externalSyncSource;
  options: {
    choices: Array<{
      id: string;
      name: string;
      color?:
        | 'blueLight2'
        | 'cyanLight2'
        | 'tealLight2'
        | 'greenLight2'
        | 'yellowLight2'
        | 'orangeLight2'
        | 'redLight2'
        | 'pinkLight2'
        | 'purpleLight2'
        | 'grayLight2'
        | 'blueLight1'
        | 'cyanLight1'
        | 'tealLight1'
        | 'greenLight1'
        | 'yellowLight1'
        | 'orangeLight1'
        | 'redLight1'
        | 'pinkLight1'
        | 'purpleLight1'
        | 'grayLight1'
        | 'blueBright'
        | 'cyanBright'
        | 'tealBright'
        | 'greenBright'
        | 'yellowBright'
        | 'orangeBright'
        | 'redBright'
        | 'pinkBright'
        | 'purpleBright'
        | 'grayBright'
        | 'blueDark1'
        | 'cyanDark1'
        | 'tealDark1'
        | 'greenDark1'
        | 'yellowDark1'
        | 'orangeDark1'
        | 'redDark1'
        | 'pinkDark1'
        | 'purpleDark1'
        | 'grayDark1';
    }>;
  };
};

export type APIFieldFormula = APIBaseField & {
  type: APIFieldType.formula;
  options: {
    isValid: boolean;
    referencedFieldIds: Array<string>;
    // `result` contains the type and field options of the evaluated field type, or null if the formula is invalid.
    result: APIField | null;
  };
};

export type APILastModifiedTime = APIBaseField & {
  type: APIFieldType.lastModifiedTime;
  options: {
    isValid: boolean;
    referencedFieldIds: Array<string>;
    result:
      | {
          type: APIFieldType.date;
          options: APIDateOptions;
        }
      | {
          type: APIFieldType.dateTime;
          options: APIDateTimeOptions;
        };
  };
};

export type APIMultipleAttachments = APIBaseField & {
  type: APIFieldType.multipleAttachments;
  options: {
    // Whether attachments are rendered in the reverse order from the cell value in the Airtable UI (i.e. most recent first). You generally do not need to rely on this option.
    isReversed: boolean;
  };
};

export type APIMultipleLookupValues = APIBaseField & {
  type: APIFieldType.multipleLookupValues;
  options: {
    isValid: boolean;
    // The linked field id
    recordLinkFieldId: string;
    // The id of the field in the linked table
    fieldIdInLinkedTable: string;
    // `result` contains the type and field options of the evaluated field type, or null if the lookup is invalid.
    result: unknown;
  };
};

export type APIMultipleRecordLinks = APIBaseField & {
  type: APIFieldType.multipleRecordLinks;
  options: {
    // Whether attachments are rendered in the reverse order from the cell value in the Airtable UI (i.e. most recent first). You generally do not need to rely on this option.
    isReversed: boolean;
    inverseLinkFieldId?: string;
    linkedTableId: string;
    viewIdForRecordSelection?: string;
    // Whether this field prefers to only have a single linked record. While this preference is enforced in the Airtable UI, it is possible for a field that prefers single linked records to have multiple record links (for example, via copy-and-paste or programmatic updates).
    prefersSingleRecordLink: boolean;
  };
};

export type APIMultipleSelects = APIBaseField & {
  type: APIFieldType.multipleSelects;
  options: {
    choices: Array<{
      id: string;
      name: string;
      color?:
        | 'blueLight2'
        | 'cyanLight2'
        | 'tealLight2'
        | 'greenLight2'
        | 'yellowLight2'
        | 'orangeLight2'
        | 'redLight2'
        | 'pinkLight2'
        | 'purpleLight2'
        | 'grayLight2'
        | 'blueLight1'
        | 'cyanLight1'
        | 'tealLight1'
        | 'greenLight1'
        | 'yellowLight1'
        | 'orangeLight1'
        | 'redLight1'
        | 'pinkLight1'
        | 'purpleLight1'
        | 'grayLight1'
        | 'blueBright'
        | 'cyanBright'
        | 'tealBright'
        | 'greenBright'
        | 'yellowBright'
        | 'orangeBright'
        | 'redBright'
        | 'pinkBright'
        | 'purpleBright'
        | 'grayBright'
        | 'blueDark1'
        | 'cyanDark1'
        | 'tealDark1'
        | 'greenDark1'
        | 'yellowDark1'
        | 'orangeDark1'
        | 'redDark1'
        | 'pinkDark1'
        | 'purpleDark1'
        | 'grayDark1';
    }>;
  };
};

export type APINumber = APIBaseField & {
  type: APIFieldType.number;
  options: {
    // from 0 to 8 inclusive
    precision: number;
  };
};

export type APIPercent = APIBaseField & {
  type: APIFieldType.percent;
  options: {
    // from 0 to 8 inclusive
    precision: number;
  };
};

export type APIRating = APIBaseField & {
  type: APIFieldType.rating;
  options: {
    color:
      | 'yellowBright'
      | 'orangeBright'
      | 'redBright'
      | 'pinkBright'
      | 'purpleBright'
      | 'blueBright'
      | 'cyanBright'
      | 'tealBright'
      | 'greenBright'
      | 'grayBright';
    icon: 'star' | 'heart' | 'thumbsUp' | 'flag' | 'dot';
    // from 1 to 10 inclusive
    max: number;
  };
};

export type APIRollup = APIBaseField & {
  type: APIFieldType.rollup;
  options: {
    isValid: boolean;
    // The linked field id
    recordLinkFieldId: string;
    // The id of the field in the linked table
    fieldIdInLinkedTable: string;
    // The ids of any fields referenced in the rollup formula
    referencedFieldIds: Array<string>;
    // `result` contains the type and field options of the evaluated field type.
    result: unknown;
  };
};

export type APIFieldSingleSelect = APIBaseField & {
  type: APIFieldType.singleSelect;
  options: {
    choices: Array<{
      id: string;
      name: string;
      color?:
        | 'blueLight2'
        | 'cyanLight2'
        | 'tealLight2'
        | 'greenLight2'
        | 'yellowLight2'
        | 'orangeLight2'
        | 'redLight2'
        | 'pinkLight2'
        | 'purpleLight2'
        | 'grayLight2'
        | 'blueLight1'
        | 'cyanLight1'
        | 'tealLight1'
        | 'greenLight1'
        | 'yellowLight1'
        | 'orangeLight1'
        | 'redLight1'
        | 'pinkLight1'
        | 'purpleLight1'
        | 'grayLight1'
        | 'blueBright'
        | 'cyanBright'
        | 'tealBright'
        | 'greenBright'
        | 'yellowBright'
        | 'orangeBright'
        | 'redBright'
        | 'pinkBright'
        | 'purpleBright'
        | 'grayBright'
        | 'blueDark1'
        | 'cyanDark1'
        | 'tealDark1'
        | 'greenDark1'
        | 'yellowDark1'
        | 'orangeDark1'
        | 'redDark1'
        | 'pinkDark1'
        | 'purpleDark1'
        | 'grayDark1';
    }>;
  };
};

export type APIField =
  | APIFieldNoOptions
  | APIFieldCheckbox
  | APIFieldCount
  | APIFieldCreatedTime
  | APIFieldCurrency
  | APIDate
  | APIDateTime
  | APIDuration
  | APIExternalSyncSource
  | APIFieldFormula
  | APILastModifiedTime
  | APIMultipleAttachments
  | APIMultipleLookupValues
  | APIMultipleRecordLinks
  | APIMultipleSelects
  | APINumber
  | APIPercent
  | APIRating
  | APIRollup
  | APIFieldSingleSelect;

export type APITable = {
  id: string;
  name: string;
  description: string;
  primaryFieldId: string;
  fields: APIField[];
};
