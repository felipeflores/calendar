import { lazy } from "react";
import { Navigate } from "react-router-dom";

const FullLayout = lazy(() => import("../layouts/FullLayout"));
const Config = lazy(() => import("../screens/Config"));

const ThemeRoutes = [
    {
        path: "/",
        element: <FullLayout />,
        children: [
            { path: "/", element: <Navigate to="/config" /> },
            { path: "/config", element: <Config /> },
        ]
    }
];

export default ThemeRoutes;