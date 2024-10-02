import App from './App';

import { GCPConnectorProvider } from './context/GCPConnectorContext';

function AppWrapper() {
  return (
    <GCPConnectorProvider>
      <App />
    </GCPConnectorProvider>
  );
}

export default AppWrapper;
