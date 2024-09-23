import { queryOptions, useQuery } from "@tanstack/react-query";
import { QueryConfig } from "../../../lib/react-query";

import {GetTodoByIdService, type GetTodosByIdData} from "../../../openapi/requests";
import { useGetTodoByIdServiceGetTodosByIdKey } from "../../../openapi/queries";

export const getTodoQueryOptions = (todoId: string) => {
    const data: GetTodosByIdData = {id: todoId};
    return queryOptions({
        queryKey: [useGetTodoByIdServiceGetTodosByIdKey, data],
        queryFn: () => {
            return GetTodoByIdService.getTodosById(data);
        },
    });
}

type UseTodoOptions = {
    todoId: string;
    queryConfig?: QueryConfig<typeof getTodoQueryOptions>;
}

export const useTodo = ({ todoId, queryConfig }: UseTodoOptions) => {
    return useQuery({
        ...getTodoQueryOptions(todoId),
        ...queryConfig,
    });
}
