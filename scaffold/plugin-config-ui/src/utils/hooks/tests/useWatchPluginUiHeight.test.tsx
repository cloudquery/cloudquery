import { render } from '@testing-library/react';
import { useWatchPluginUiHeight } from '../useWatchPluginUiHeight';

// global.ResizeObserver = require('resize-observer-polyfill');

const App = () => {
  const containerRef = useWatchPluginUiHeight();

  return (
    <div ref={containerRef}>
      <h1>Test</h1>
    </div>
  );
};

let originalPostMessage: typeof window.postMessage;

beforeEach(() => {
  originalPostMessage = window.parent.postMessage;
  window.parent.postMessage = jest.fn();
  jest.spyOn(window, 'addEventListener');
});

afterEach(() => {
  window.parent.postMessage = originalPostMessage;
  jest.restoreAllMocks();
});

test('useWatchPluginUiHeight', async () => {
  render(<App />);

  expect(window.parent.postMessage).toBeCalledWith(
    {
      type: 'height_changed',
      payload: { height: 0 },
      id: expect.any(String),
    },
    '*',
  );
});
