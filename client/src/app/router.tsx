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
                const { TodosRoute } = await import('./routes/todos/todos');
                return {Component: TodosRoute};
            },
            loader: async (args: LoaderFunctionArgs) => {
                const { TodosLoader } = await import('./routes/todos/todos');
                return TodosLoader(queryClient)(args);
            }
        },
        {
            path: '/todos/:todoId',
            lazy: async () => {
                const { TodoRoute } = await import('./routes/todos/todo');
                return {Component: TodoRoute};
            },
            loader: async (args: LoaderFunctionArgs) => {
                const { TodoLoader } = await import('./routes/todos/todo');
                return TodoLoader(queryClient)(args);
            }
        }
    ])


export const AppRouter = () => {
    const queryClient = useQueryClient();
    const router = createAppRouter(queryClient);
    return <RouterProvider router={router} />;
}
