import { Suspense } from "react";
import { ErrorBoundary } from "react-error-boundary";
import { Outlet, useLocation } from "react-router-dom";

export const AppRoot = () => {
    const location = useLocation();
    return (
       <Suspense
           fallback={
               <div className="flex size-full items-center justify-center">
                   Loading...
               </div>
           }
       >
           <ErrorBoundary
               key={location.pathname}
               fallback={<div>Something went wrong</div>}
           >
                <Outlet />
           </ErrorBoundary>
       </Suspense>
    );
}
