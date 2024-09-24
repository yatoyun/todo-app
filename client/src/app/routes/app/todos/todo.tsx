import { QueryClient } from "@tanstack/react-query";
import { useParams, LoaderFunctionArgs } from "react-router-dom";

import { getTodoQueryOptions, useTodo } from "../../../../features/todos/api/get-todo";

export const TodoLoader = (queryClient: QueryClient) =>
    async ({ params }: LoaderFunctionArgs) => {
        const todoId = params.todoId as string;
        const todoQuery = getTodoQueryOptions(todoId);

        const promises = [
            queryClient.getQueryData(todoQuery.queryKey) ??
            (await queryClient.fetchQuery(todoQuery)),
        ] as const;

        const [todo] = await Promise.all(promises);
        return {
            todo,
        };
    };

export const TodoRoute = () => {
    const params = useParams();
    const todoId = params.todoId as string;
    const todoQuery = useTodo({todoId});

    if (todoQuery.isLoading) {
        return <div>Loading...</div>;
    }

    const todo = todoQuery.data;
    if (!todo) return null;

    return (
        <div>
            <h1>{todo.title}</h1>
            <p>{todo.description}</p>
        </div>
    );
}
