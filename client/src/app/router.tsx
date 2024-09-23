import { QueryClient, useQueryClient } from "@tanstack/react-query";
import { createBrowserRouter, RouterProvider, LoaderFunctionArgs } from 'react-router-dom';

export const createAppRouter = (queryClient: QueryClient) =>
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
                const { TodosRoute } = await import('./routes/todos');
                return {Component: TodosRoute};
            },
            loader: async (args: LoaderFunctionArgs) => {
                const { TodosLoader } = await import('./routes/todos');
                return TodosLoader(queryClient)(args);
            }
        }
    ])


export const AppRouter = () => {
    const queryClient = useQueryClient();
    const router = createAppRouter(queryClient);
    return <RouterProvider router={router} />;
}
