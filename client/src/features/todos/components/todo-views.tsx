import { useTodo } from "../api/get-todo";
import { formatDateTime } from "../../../utils/format";

export const TodoView = ({ todoId }: {todoId: string}) => {
    const todoQuery = useTodo({ todoId });

    if (todoQuery.isLoading) {
        return <div>Loading...</div>;
    }

    const todo = todoQuery?.data;
    if (!todo) return null;

    return (
        <div>
            <span className="text-ts font-bold">
                {formatDateTime(todo.created_at ?? 'Timestamp not found')}
            </span>
            <div className="mt-6 flex flex-col space-y-16">
                <div>
                    <div className="overflow-hidden bg-white shadow sm:rounded-lg">
                        <div className="px-4 py-5 sm:px-6">
                            <h3 className="text-lg font-medium leading-6 text-gray-900">
                                {todo.title}
                            </h3>
                            <div className="mt-1 max-w-2xl text-sm text-gray-500">
                                {todo.description}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};
