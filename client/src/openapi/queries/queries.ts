// generated with @7nohe/openapi-react-query-codegen@1.6.1 

import { useMutation, UseMutationOptions, useQuery, UseQueryOptions } from "@tanstack/react-query";
import { CreateTodoService, CreateUserService, DeleteTodoService, DeleteUserService, GetTodoByIdService, GetTodosService, GetUserByAuth0IdService, GetUserByIdService, GetUsersService, UpdateTodoService, UpdateUserService } from "../requests/services.gen";
import { usecase_CreateTodoRequest, usecase_CreateUserRequest, usecase_UpdateTodoRequest, usecase_UpdateUserRequest } from "../requests/types.gen";
import * as Common from "./common";
/**
* todo一覧を取得
* @returns usecase_TodoResponse OK
* @throws ApiError
*/
export const useGetTodosServiceGetTodos = <TData = Common.GetTodosServiceGetTodosDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>(queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useQuery<TData, TError>({ queryKey: Common.UseGetTodosServiceGetTodosKeyFn(queryKey), queryFn: () => GetTodosService.getTodos() as TData, ...options });
/**
* todoを取得
* @param data The data for the request.
* @param data.id Todo ID
* @returns entity_Todo OK
* @throws ApiError
*/
export const useGetTodoByIdServiceGetTodosById = <TData = Common.GetTodoByIdServiceGetTodosByIdDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>({ id }: {
  id: string;
}, queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useQuery<TData, TError>({ queryKey: Common.UseGetTodoByIdServiceGetTodosByIdKeyFn({ id }, queryKey), queryFn: () => GetTodoByIdService.getTodosById({ id }) as TData, ...options });
/**
* user一覧を取得する
* user一覧を取得する
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const useGetUsersServiceGetUsers = <TData = Common.GetUsersServiceGetUsersDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>(queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useQuery<TData, TError>({ queryKey: Common.UseGetUsersServiceGetUsersKeyFn(queryKey), queryFn: () => GetUsersService.getUsers() as TData, ...options });
/**
* userをIDで取得する
* userをIDで取得する
* @param data The data for the request.
* @param data.id ID
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const useGetUserByIdServiceGetUsersById = <TData = Common.GetUserByIdServiceGetUsersByIdDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>({ id }: {
  id: string;
}, queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useQuery<TData, TError>({ queryKey: Common.UseGetUserByIdServiceGetUsersByIdKeyFn({ id }, queryKey), queryFn: () => GetUserByIdService.getUsersById({ id }) as TData, ...options });
/**
* userをAuth0IDで取得する
* userをAuth0IDで取得する
* @param data The data for the request.
* @param data.auth0Id Auth0ID
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const useGetUserByAuth0IdServiceGetUsersAuth0ByAuth0Id = <TData = Common.GetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>({ auth0Id }: {
  auth0Id: string;
}, queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useQuery<TData, TError>({ queryKey: Common.UseGetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdKeyFn({ auth0Id }, queryKey), queryFn: () => GetUserByAuth0IdService.getUsersAuth0ByAuth0Id({ auth0Id }) as TData, ...options });
/**
* todoを作成
* @param data The data for the request.
* @param data.requestBody Todo
* @returns usecase_TodoResponse Created
* @throws ApiError
*/
export const useCreateTodoServicePostTodos = <TData = Common.CreateTodoServicePostTodosMutationResult, TError = unknown, TContext = unknown>(options?: Omit<UseMutationOptions<TData, TError, {
  requestBody: usecase_CreateTodoRequest;
}, TContext>, "mutationFn">) => useMutation<TData, TError, {
  requestBody: usecase_CreateTodoRequest;
}, TContext>({ mutationFn: ({ requestBody }) => CreateTodoService.postTodos({ requestBody }) as unknown as Promise<TData>, ...options });
/**
* userを作成する
* userを作成する
* @param data The data for the request.
* @param data.requestBody user
* @returns usecase_UserResponse Created
* @throws ApiError
*/
export const useCreateUserServicePostUsers = <TData = Common.CreateUserServicePostUsersMutationResult, TError = unknown, TContext = unknown>(options?: Omit<UseMutationOptions<TData, TError, {
  requestBody: usecase_CreateUserRequest;
}, TContext>, "mutationFn">) => useMutation<TData, TError, {
  requestBody: usecase_CreateUserRequest;
}, TContext>({ mutationFn: ({ requestBody }) => CreateUserService.postUsers({ requestBody }) as unknown as Promise<TData>, ...options });
/**
* todoを更新
* @param data The data for the request.
* @param data.id Todo ID
* @param data.requestBody Todo
* @returns usecase_TodoResponse OK
* @throws ApiError
*/
export const useUpdateTodoServicePutTodosById = <TData = Common.UpdateTodoServicePutTodosByIdMutationResult, TError = unknown, TContext = unknown>(options?: Omit<UseMutationOptions<TData, TError, {
  id: string;
  requestBody: usecase_UpdateTodoRequest;
}, TContext>, "mutationFn">) => useMutation<TData, TError, {
  id: string;
  requestBody: usecase_UpdateTodoRequest;
}, TContext>({ mutationFn: ({ id, requestBody }) => UpdateTodoService.putTodosById({ id, requestBody }) as unknown as Promise<TData>, ...options });
/**
* userを更新する
* userを更新する
* @param data The data for the request.
* @param data.id ID
* @param data.requestBody user
* @returns string User updated successfully
* @throws ApiError
*/
export const useUpdateUserServicePutUsersById = <TData = Common.UpdateUserServicePutUsersByIdMutationResult, TError = unknown, TContext = unknown>(options?: Omit<UseMutationOptions<TData, TError, {
  id: string;
  requestBody: usecase_UpdateUserRequest;
}, TContext>, "mutationFn">) => useMutation<TData, TError, {
  id: string;
  requestBody: usecase_UpdateUserRequest;
}, TContext>({ mutationFn: ({ id, requestBody }) => UpdateUserService.putUsersById({ id, requestBody }) as unknown as Promise<TData>, ...options });
/**
* todoを削除
* @param data The data for the request.
* @param data.id Todo ID
* @returns void No Content
* @throws ApiError
*/
export const useDeleteTodoServiceDeleteTodosById = <TData = Common.DeleteTodoServiceDeleteTodosByIdMutationResult, TError = unknown, TContext = unknown>(options?: Omit<UseMutationOptions<TData, TError, {
  id: string;
}, TContext>, "mutationFn">) => useMutation<TData, TError, {
  id: string;
}, TContext>({ mutationFn: ({ id }) => DeleteTodoService.deleteTodosById({ id }) as unknown as Promise<TData>, ...options });
/**
* userを削除する
* userを削除する
* @param data The data for the request.
* @param data.id ID
* @returns string User deleted successfully
* @throws ApiError
*/
export const useDeleteUserServiceDeleteUsersById = <TData = Common.DeleteUserServiceDeleteUsersByIdMutationResult, TError = unknown, TContext = unknown>(options?: Omit<UseMutationOptions<TData, TError, {
  id: string;
}, TContext>, "mutationFn">) => useMutation<TData, TError, {
  id: string;
}, TContext>({ mutationFn: ({ id }) => DeleteUserService.deleteUsersById({ id }) as unknown as Promise<TData>, ...options });
