import React from 'react';

import { AppProvider } from "./provider";
import { AppRouter } from './router';

// FIXME
// add provider
const App: React.FC = () => {
    return (
        <AppProvider>
            <AppRouter />
        </AppProvider>
    );
}

export default App;
