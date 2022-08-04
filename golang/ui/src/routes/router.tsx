import { lazy } from "react";
import { Navigate } from "react-router-dom";

const FullLayout = lazy(() => import("../layouts/FullLayout"));
const Config = lazy(() => import("../screens/Config"));
const Start = lazy(() => import("../screens/Start"));
const Ports = lazy(() => import("../screens/Ports"));
const Networks = lazy(() => import("../screens/Networks"));
const Info = lazy(() => import("../screens/Info"));
const Google = lazy(() => import("../screens/google/Google"));
const Auth = lazy(() => import("../screens/google/Auth"));
const Onboard = lazy(() => import("../screens/Onboard"));

const ThemeRoutes = [
    {
        path: "/",
        element: <FullLayout />,
        children: [
            { 
                path: "", element: <Onboard />
            },
            { 
                path: "teste", element: <Auth />
            },
            {
                path: "config",
                element: <Config />,
                children: [
                    { 
                        children: [
                            { path: "", element: <Start />, },
                            { path: "google/credentials", element: <Google />, },
                            { path: "google/auth", element: <Auth />, },
                            { path: "ports", element: <Ports />, },
                            { path: "info", element: <Info />, },
                            { path: "networks", element: <Networks />, }
                        ]
                    },
                ]
            }
        ]
    },    
];

export default ThemeRoutes;