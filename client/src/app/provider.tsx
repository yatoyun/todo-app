import React, { useEffect } from 'react';
import { Auth0Provider, useAuth0 } from "@auth0/auth0-react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';

import { queryConfig } from '../lib/react-query';
import { auth0Config } from "../config/auth0";
import { setGetToken } from "../lib/tokenProvider";

interface  Auth0SetterProps {
    children: React.ReactNode;
}

type AppProviderProps = {
  children: React.ReactNode;
};

/**
 * Auth0 のアクセストークン取得関数をトークンプロバイダーに設定します。
 */
const Auth0Setter = ({ children }: Auth0SetterProps) => {
    const { getAccessTokenSilently } = useAuth0();
    useEffect(() => {
        setGetToken(() => getAccessTokenSilently());
    }, [getAccessTokenSilently]);

    return <>{children}</>;
};

/**
 * アプリケーション全体のプロバイダーを設定します。
 * - Auth0Provider
 * - Auth0Setter
 * - QueryClientProvider
 */
export const AppProvider = ({ children }: AppProviderProps) => {
    const [queryClient] = React.useState(
        () =>
            new QueryClient({
                defaultOptions: queryConfig,
            }),
    );

    return (
        <React.Suspense
          fallback={
            <div className="flex h-screen w-screen items-center justify-center">
                <div>Loading...</div>
            </div>
          }
        >
            <Auth0Provider
                domain={auth0Config.domain}
                clientId={auth0Config.clientId}
                authorizationParams={{
                    redirect_uri: auth0Config.redirectUri,
                    audience: auth0Config.audience
                }}
            >
                <Auth0Setter>
                    <QueryClientProvider client={queryClient}>
                        {<ReactQueryDevtools/>}
                        {children}
                    </QueryClientProvider>
                </Auth0Setter>
            </Auth0Provider>
        </React.Suspense>
    );
};
