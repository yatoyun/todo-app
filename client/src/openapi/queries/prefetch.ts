// generated with @7nohe/openapi-react-query-codegen@1.6.1 

import { type QueryClient } from "@tanstack/react-query";
import { GetTodoByIdService, GetTodosService, GetUserByAuth0IdService, GetUserByIdService, GetUsersService } from "../requests/services.gen";
import * as Common from "./common";
/**
* todo一覧を取得
* @returns usecase_TodoResponse OK
* @throws ApiError
*/
export const prefetchUseGetTodosServiceGetTodos = (queryClient: QueryClient) => queryClient.prefetchQuery({ queryKey: Common.UseGetTodosServiceGetTodosKeyFn(), queryFn: () => GetTodosService.getTodos() });
/**
* todoを取得
* @param data The data for the request.
* @param data.id Todo ID
* @returns entity_Todo OK
* @throws ApiError
*/
export const prefetchUseGetTodoByIdServiceGetTodosById = (queryClient: QueryClient, { id }: {
  id: string;
}) => queryClient.prefetchQuery({ queryKey: Common.UseGetTodoByIdServiceGetTodosByIdKeyFn({ id }), queryFn: () => GetTodoByIdService.getTodosById({ id }) });
/**
* user一覧を取得する
* user一覧を取得する
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const prefetchUseGetUsersServiceGetUsers = (queryClient: QueryClient) => queryClient.prefetchQuery({ queryKey: Common.UseGetUsersServiceGetUsersKeyFn(), queryFn: () => GetUsersService.getUsers() });
/**
* userをIDで取得する
* userをIDで取得する
* @param data The data for the request.
* @param data.id ID
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const prefetchUseGetUserByIdServiceGetUsersById = (queryClient: QueryClient, { id }: {
  id: string;
}) => queryClient.prefetchQuery({ queryKey: Common.UseGetUserByIdServiceGetUsersByIdKeyFn({ id }), queryFn: () => GetUserByIdService.getUsersById({ id }) });
/**
* userをAuth0IDで取得する
* userをAuth0IDで取得する
* @param data The data for the request.
* @param data.auth0Id Auth0ID
* @returns usecase_UserResponse OK
* @throws ApiError
*/
export const prefetchUseGetUserByAuth0IdServiceGetUsersAuth0ByAuth0Id = (queryClient: QueryClient, { auth0Id }: {
  auth0Id: string;
}) => queryClient.prefetchQuery({ queryKey: Common.UseGetUserByAuth0IdServiceGetUsersAuth0ByAuth0IdKeyFn({ auth0Id }), queryFn: () => GetUserByAuth0IdService.getUsersAuth0ByAuth0Id({ auth0Id }) });
