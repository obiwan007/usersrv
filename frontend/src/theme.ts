import { deepOrange, red, yellow } from "@material-ui/core/colors";
import { createMuiTheme, responsiveFontSizes } from "@material-ui/core/styles";
// A custom theme for this app
let theme = createMuiTheme({
  typography: {
    // In Chinese and Japanese the characters are usually larger,
    // so a smaller fontsize may be appropriate.
    // fontSize: 12,
  },
  palette: {
    type: "dark",
    primary: deepOrange,

    secondary: yellow,

    error: {
      main: red.A400,
    },
  },
});

theme = responsiveFontSizes(theme);

export default theme;
