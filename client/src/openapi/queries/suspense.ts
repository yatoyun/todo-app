// generated with @7nohe/openapi-react-query-codegen@1.6.1 

import { UseQueryOptions, useSuspenseQuery } from "@tanstack/react-query";
import { GetTodoByIdService, GetTodosService, GetUserByAuth0IdService, GetUserByIdService, GetUsersService } from "../requests/services.gen";
import * as Common from "./common";
/**
* todo一覧を取得
* @returns usecase_TodoResponse OK
* @throws ApiError
*/
export const useGetTodosServiceGetTodosSuspense = <TData = Common.GetTodosServiceGetTodosDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>(queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useSuspenseQuery<TData, TError>({ queryKey: Common.UseGetTodosServiceGetTodosKeyFn(queryKey), queryFn: () => GetTodosService.getTodos() as TData, ...options });
/**
* todoを取得
* @param data The data for the request.
* @param data.id Todo ID
* @returns entity_Todo OK
* @throws ApiError
*/
export const useGetTodoByIdServiceGetTodosByIdSuspense = <TData = Common.GetTodoByIdServiceGetTodosByIdDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>({ id }: {
  id: string;
}, queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useSuspenseQuery<TData, TError>({ queryKey: Common.UseGetTodoByIdServiceGetTodosByIdKeyFn({ id }, queryKey), queryFn: () => GetTodoByIdService.getTodosById({ id }) as TData, ...options });
/**
* user一覧を取得する
* user一覧を取得する
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const useGetUsersServiceGetUsersSuspense = <TData = Common.GetUsersServiceGetUsersDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>(queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useSuspenseQuery<TData, TError>({ queryKey: Common.UseGetUsersServiceGetUsersKeyFn(queryKey), queryFn: () => GetUsersService.getUsers() as TData, ...options });
/**
* userをIDで取得する
* userをIDで取得する
* @param data The data for the request.
* @param data.id ID
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const useGetUserByIdServiceGetUsersByIdSuspense = <TData = Common.GetUserByIdServiceGetUsersByIdDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>({ id }: {
  id: string;
}, queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useSuspenseQuery<TData, TError>({ queryKey: Common.UseGetUserByIdServiceGetUsersByIdKeyFn({ id }, queryKey), queryFn: () => GetUserByIdService.getUsersById({ id }) as TData, ...options });
/**
* userをAuth0IDで取得する
* userをAuth0IDで取得する
* @param data The data for the request.
* @param data.auth0Id Auth0ID
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const useGetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdSuspense = <TData = Common.GetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdDefaultResponse, TError = unknown, TQueryKey extends Array<unknown> = unknown[]>({ auth0Id }: {
  auth0Id: string;
}, queryKey?: TQueryKey, options?: Omit<UseQueryOptions<TData, TError>, "queryKey" | "queryFn">) => useSuspenseQuery<TData, TError>({ queryKey: Common.UseGetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdKeyFn({ auth0Id }, queryKey), queryFn: () => GetUserByAuth0IdService.getUsersAuth0ByAuth0Id({ auth0Id }) as TData, ...options });
