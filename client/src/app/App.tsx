import React from 'react';
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import {AppRouter} from './router';
import { queryConfig } from '../lib/react-query';

// FIXME
// add provider
const App: React.FC = () => {
    const [queryClient] = React.useState(
        () =>
          new QueryClient({
            defaultOptions: queryConfig,
          }),
      );
    return (
        <QueryClientProvider client={queryClient}>
            <AppRouter />
        </QueryClientProvider>
    );
}

export default App;
