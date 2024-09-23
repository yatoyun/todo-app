import { useQueryClient } from "@tanstack/react-query";
import { useSearchParams } from "react-router-dom";

import  { getTodosQueryOptions } from "../api/get-todos";
import { useTodos } from "../api/get-todos";

export const TodoList = () => {
    const searchParams = useSearchParams();
    const todoQuery = useTodos();
    const queryClient = useQueryClient();
    if (todoQuery.isLoading) {
        return <div>Loading...</div>;
    }
    const todos = todoQuery.data;
    // log the todos
    console.log(todos);
    if (!todos) return null;
    return (
        <div className="mt-4">
            <h2 className="text-2xl font-bold">Todos</h2>
            <ul>
                {todos.map((todo) => (
                    <li key={todo.id}>
                            {todo.title}
                    </li>
                ))}
            </ul>
        </div>
    );
}
