import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from './routes/index';
// import Todos from './routes/todos';

const AppRouter: React.FC = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            {/*<Route path="/todos" element={<Todos />} />*/}
        </Routes>
    );
}
export default AppRouter;
