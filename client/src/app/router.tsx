import {createBrowserRouter, RouterProvider} from 'react-router-dom';

// FIXME
// add query client
export const createAppRouter = () =>
    createBrowserRouter([
        {
            path: '/',
            lazy: async () => {
                const { default: Home } = await import('./routes/index');
                return {Component: Home};
            }
        },
        {
            path: '/todos',
            lazy: async () => {
                const { default: Todos } = await import('./routes/todos');
                return {Component: Todos};
            }
        }
    ])


export const AppRouter = () => {
    const router = createAppRouter();
    return <RouterProvider router={router} />;
}
