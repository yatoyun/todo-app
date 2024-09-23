import { queryOptions, useQuery } from '@tanstack/react-query';
import { QueryConfig } from '../../../lib/react-query';

import { GetTodosService } from '../../../openapi/requests';
import { useGetTodoByIdServiceGetTodosByIdKey } from '../../../openapi/queries';

export const getTodosQueryOptions = () => {
    return queryOptions({
        queryKey: [useGetTodoByIdServiceGetTodosByIdKey],
        queryFn: () => {
            return GetTodosService.getTodos();
        },
    });
}

type UseTodosOptions = {
    queryConfig?: QueryConfig<typeof getTodosQueryOptions>;
}

export const useTodos = ({ queryConfig }: UseTodosOptions = {}) => {
    return useQuery({
        ...getTodosQueryOptions(),
        ...queryConfig,
    });
}
