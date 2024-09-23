// generated with @7nohe/openapi-react-query-codegen@1.6.1 

import { UseQueryResult } from "@tanstack/react-query";
import { CreateTodoService, CreateUserService, DeleteTodoService, DeleteUserService, GetTodoByIdService, GetTodosService, GetUserByAuth0IdService, GetUserByIdService, GetUsersService, UpdateTodoService, UpdateUserService } from "../requests/services.gen";
export type GetTodosServiceGetTodosDefaultResponse = Awaited<ReturnType<typeof GetTodosService.getTodos>>;
export type GetTodosServiceGetTodosQueryResult<TData = GetTodosServiceGetTodosDefaultResponse, TError = unknown> = UseQueryResult<TData, TError>;
export const useGetTodosServiceGetTodosKey = "GetTodosServiceGetTodos";
export const UseGetTodosServiceGetTodosKeyFn = (queryKey?: Array<unknown>) => [useGetTodosServiceGetTodosKey, ...(queryKey ?? [])];
export type GetTodoByIdServiceGetTodosByIdDefaultResponse = Awaited<ReturnType<typeof GetTodoByIdService.getTodosById>>;
export type GetTodoByIdServiceGetTodosByIdQueryResult<TData = GetTodoByIdServiceGetTodosByIdDefaultResponse, TError = unknown> = UseQueryResult<TData, TError>;
export const useGetTodoByIdServiceGetTodosByIdKey = "GetTodoByIdServiceGetTodosById";
export const UseGetTodoByIdServiceGetTodosByIdKeyFn = ({ id }: {
  id: string;
}, queryKey?: Array<unknown>) => [useGetTodoByIdServiceGetTodosByIdKey, ...(queryKey ?? [{ id }])];
export type GetUsersServiceGetUsersDefaultResponse = Awaited<ReturnType<typeof GetUsersService.getUsers>>;
export type GetUsersServiceGetUsersQueryResult<TData = GetUsersServiceGetUsersDefaultResponse, TError = unknown> = UseQueryResult<TData, TError>;
export const useGetUsersServiceGetUsersKey = "GetUsersServiceGetUsers";
export const UseGetUsersServiceGetUsersKeyFn = (queryKey?: Array<unknown>) => [useGetUsersServiceGetUsersKey, ...(queryKey ?? [])];
export type GetUserByIdServiceGetUsersByIdDefaultResponse = Awaited<ReturnType<typeof GetUserByIdService.getUsersById>>;
export type GetUserByIdServiceGetUsersByIdQueryResult<TData = GetUserByIdServiceGetUsersByIdDefaultResponse, TError = unknown> = UseQueryResult<TData, TError>;
export const useGetUserByIdServiceGetUsersByIdKey = "GetUserByIdServiceGetUsersById";
export const UseGetUserByIdServiceGetUsersByIdKeyFn = ({ id }: {
  id: string;
}, queryKey?: Array<unknown>) => [useGetUserByIdServiceGetUsersByIdKey, ...(queryKey ?? [{ id }])];
export type GetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdDefaultResponse = Awaited<ReturnType<typeof GetUserByAuth0IdService.getUsersAuth0ByAuth0Id>>;
export type GetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdQueryResult<TData = GetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdDefaultResponse, TError = unknown> = UseQueryResult<TData, TError>;
export const useGetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdKey = "GetUserByAuth0IdServiceGetUsersAuth0ByAuth0Id";
export const UseGetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdKeyFn = ({ auth0Id }: {
  auth0Id: string;
}, queryKey?: Array<unknown>) => [useGetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdKey, ...(queryKey ?? [{ auth0Id }])];
export type CreateTodoServicePostTodosMutationResult = Awaited<ReturnType<typeof CreateTodoService.postTodos>>;
export type CreateUserServicePostUsersMutationResult = Awaited<ReturnType<typeof CreateUserService.postUsers>>;
export type UpdateTodoServicePutTodosByIdMutationResult = Awaited<ReturnType<typeof UpdateTodoService.putTodosById>>;
export type UpdateUserServicePutUsersByIdMutationResult = Awaited<ReturnType<typeof UpdateUserService.putUsersById>>;
export type DeleteTodoServiceDeleteTodosByIdMutationResult = Awaited<ReturnType<typeof DeleteTodoService.deleteTodosById>>;
export type DeleteUserServiceDeleteUsersByIdMutationResult = Awaited<ReturnType<typeof DeleteUserService.deleteUsersById>>;
