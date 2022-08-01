import { lazy } from "react";
import { Navigate } from "react-router-dom";

const FullLayout = lazy(() => import("../layouts/FullLayout"));
const Config = lazy(() => import("../screens/Config"));
const Start = lazy(() => import("../screens/Start"));
const Ports = lazy(() => import("../screens/Ports"));
const Networks = lazy(() => import("../screens/Networks"));
const Info = lazy(() => import("../screens/Info"));

const ThemeRoutes = [
    {
        path: "/",
        element: <FullLayout />,
        children: [
            { path: "/", element: <Navigate to="/config" /> },
            { 
                path: "config", element: <Config />,
                children: [
                    { path: "", element: <Start />, },
                    { path: "ports", element: <Ports />, },
                    { path: "info", element: <Info />, },
                    { path: "networks", element: <Networks />, }
                ]
            },
        ]
    }
];

export default ThemeRoutes;