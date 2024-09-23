import { QueryClient } from '@tanstack/react-query';
import { LoaderFunctionArgs } from 'react-router-dom';

import { getTodosQueryOptions } from '../../features/todos/api/get-todos';
import { TodoList } from '../../features/todos/components/todos-list';

export const TodosLoader = (queryClient: QueryClient) =>
    async ({ request }: LoaderFunctionArgs) => {
        const url = new URL(request.url);
        const query = getTodosQueryOptions();
        return (
            queryClient.getQueryData(query.queryKey) ??
            (await queryClient.fetchQuery({
                queryKey: query.queryKey,
                queryFn: query.queryFn,
              })
            )
        );
    };

export const TodosRoute = () => {
    return (
        <div className="mt-4">
            <TodoList />
        </div>
    );
}
