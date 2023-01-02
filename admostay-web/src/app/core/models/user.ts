// User is used to wrap the firebase user information
import firebase from "firebase/compat";

export interface User {
  uid?: string;
  email: string;
  displayName: string;
  photoURL: string;
  emailVerified: boolean;
}
