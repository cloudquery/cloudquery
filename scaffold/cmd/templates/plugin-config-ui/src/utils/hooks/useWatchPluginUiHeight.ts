import { useEffect, useRef } from 'react';
import { pluginUiMessageHandler } from '../messageHandler';

export function useWatchPluginUiHeight() {
  const containerRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const observer = new ResizeObserver(() => {
      pluginUiMessageHandler.sendMessage('height_changed', {
        height: containerRef.current?.offsetHeight ?? 0,
      });
    });

    observer.observe(containerRef.current ?? document.body);

    pluginUiMessageHandler.sendMessage('height_changed', {
      height: containerRef.current?.offsetHeight ?? 0,
    });

    return () => {
      observer.disconnect();
    };
  }, []);

  return containerRef;
}
