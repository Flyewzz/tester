import { createMuiTheme } from "@material-ui/core/styles";

// see https://material-ui.com/customization/themes/
const customTheme = createMuiTheme({
  palette: {
    type: 'dark',
    primary: {
      // light: will be calculated from palette.primary.main,
      main: "#CEE8ED",
      // dark: will be calculated from palette.primary.main,
      contrastText: "#444444",
    },
    secondary: {
      main: "#006064"
      // dark: will be calculated from palette.secondary.main,
    }
    // error: will use the default color
  }
});

export default customTheme;