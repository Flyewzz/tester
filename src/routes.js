import React from "react";
import { Redirect } from "react-router-dom";

// Layout Types
import { DefaultLayout } from "./layouts";

// Route Views
import BlogOverview from "./views/BlogOverview";
import UserProfileLite from "./views/UserProfileLite";
import AddNewPost from "./views/AddNewPost";
import Errors from "./views/Errors";
import ComponentsOverview from "./views/ComponentsOverview";
import Tables from "./views/Tables";
import BlogPosts from "./views/BlogPosts";
import MainPage from "./views/MainPage";
import ProfilePage from "./views/ProfilePage";
import LoginPage from "./views/auth/LoginPage";
import SignUpPage from "./views/auth/SignUpPage";

export default [
  {
    path: "/",
    exact: true,
    layout: DefaultLayout,
    component: BlogOverview,
    component: () => <Redirect to="/1" />,
  },

  {
    path: "/blog-overview",
    layout: DefaultLayout,
    component: BlogOverview,
  },
  {
    path: "/user-profile-lite",
    layout: DefaultLayout,
    component: UserProfileLite,
  },
  {
    path: "/add-new-post",
    layout: DefaultLayout,
    component: AddNewPost,
  },
  {
    path: "/errors",
    layout: DefaultLayout,
    component: Errors,
  },
  {
    path: "/components-overview",
    layout: DefaultLayout,
    component: ComponentsOverview,
  },
  {
    path: "/tables",
    layout: DefaultLayout,
    component: Tables,
  },
  {
    path: "/blog-posts",
    layout: DefaultLayout,
    component: BlogPosts,
  },
  {
    path: "/login",
    layout: DefaultLayout,
    component: LoginPage,
  },
  {
    path: "/signup",
    layout: DefaultLayout,
    component: SignUpPage,
  },
  {
    path: "/:id",
    exact: true,
    layout: DefaultLayout,
    component: BlogOverview,
  },
  // {
  //   path: "/profile",
  //   layout: DefaultLayout,
  //   component: ProfilePage
  // },
];
