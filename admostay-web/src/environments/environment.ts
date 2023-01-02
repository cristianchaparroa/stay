// This file can be replaced during build by using the `fileReplacements` array.
// `ng build` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
  production: false,
  apiURL :"http://localhost:8080",
  firebaseConfig: {
    apiKey: "AIzaSyB9QO01D80_2TrGcl-ImU8DTfLqO8bi9wk",
    authDomain: "stay-5c81e.firebaseapp.com",
    projectId: "stay-5c81e",
    storageBucket: "stay-5c81e.appspot.com",
    messagingSenderId: "126343230844",
    appId: "1:126343230844:web:ff7c679c1b003d397660f9",
    databaseURL: "https://stay-5c81e-default-rtdb.firebaseio.com",
  }
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.
