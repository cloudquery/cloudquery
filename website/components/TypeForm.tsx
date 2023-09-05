import { Widget } from '@typeform/embed-react';

export const TypeForm = ({
                             id,
                         }: {
    id: string;
}) => {
    const widgetContainerStyle = {
        width: '100%',
        height: 600,
        margin: '20px auto',
    }
    return <Widget
        id={id}
        style={widgetContainerStyle}
        iframeProps={{ title: 'Register for CloudQuery Cloud' }}
    />
}