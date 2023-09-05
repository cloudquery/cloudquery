import { Widget } from '@typeform/embed-react';

export const Typeform = ({
                             id,
                         }: {
    id: string;
}) => {
    const widgetContainerStyle = {
        width: '100%',
        height: 600,
        margin: '20px auto',
        background: 'white',
    }
    return <Widget
        id={id}
        style={widgetContainerStyle}
        iframeProps={{ title: 'Register for CloudQuery Cloud' }}
    />
}