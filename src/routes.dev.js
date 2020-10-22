"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports["default"] = void 0;

var _react = _interopRequireDefault(require("react"));

var _reactRouterDom = require("react-router-dom");

var _layouts = require("./layouts");

var _BlogOverview = _interopRequireDefault(require("./views/BlogOverview"));

var _UserProfileLite = _interopRequireDefault(require("./views/UserProfileLite"));

var _AddNewPost = _interopRequireDefault(require("./views/AddNewPost"));

var _Errors = _interopRequireDefault(require("./views/Errors"));

var _ComponentsOverview = _interopRequireDefault(require("./views/ComponentsOverview"));

var _Tables = _interopRequireDefault(require("./views/Tables"));

var _BlogPosts = _interopRequireDefault(require("./views/BlogPosts"));

var _MainPage = _interopRequireDefault(require("./views/MainPage"));

var _ProfilePage = _interopRequireDefault(require("./views/ProfilePage"));

var _LoginPage = _interopRequireDefault(require("./views/auth/LoginPage"));

var _SignUpPage = _interopRequireDefault(require("./views/auth/SignUpPage"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

// Layout Types
// Route Views
var _default = [{
  path: "/",
  exact: true,
  layout: _layouts.DefaultLayout,
  component: _BlogOverview["default"] // component: () => <Redirect to="/1" />,

}, // {
//   path: "/:id",
//   exact: true,
//   layout: DefaultLayout,
//   component: BlogOverview,
// },
{
  path: "/blog-overview",
  layout: _layouts.DefaultLayout,
  component: _BlogOverview["default"]
}, {
  path: "/user-profile-lite",
  layout: _layouts.DefaultLayout,
  component: _UserProfileLite["default"]
}, {
  path: "/add-new-post",
  layout: _layouts.DefaultLayout,
  component: _AddNewPost["default"]
}, {
  path: "/errors",
  layout: _layouts.DefaultLayout,
  component: _Errors["default"]
}, {
  path: "/components-overview",
  layout: _layouts.DefaultLayout,
  component: _ComponentsOverview["default"]
}, {
  path: "/tables",
  layout: _layouts.DefaultLayout,
  component: _Tables["default"]
}, {
  path: "/blog-posts",
  layout: _layouts.DefaultLayout,
  component: _BlogPosts["default"]
}, {
  path: "/login",
  layout: _layouts.DefaultLayout,
  component: _LoginPage["default"]
}, {
  path: "/signup",
  layout: _layouts.DefaultLayout,
  component: _SignUpPage["default"]
} // {
//   path: "/profile",
//   layout: DefaultLayout,
//   component: ProfilePage
// },
];
exports["default"] = _default;