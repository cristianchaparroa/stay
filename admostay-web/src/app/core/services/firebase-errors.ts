import {Injectable} from "@angular/core";

@Injectable({
  providedIn: 'root',
})
export class FirebaseErrorService {

   constructor() {}

   public GetAuthError(error: { code: string; message: "" }) {
    if (error.code === "auth/weak-password") {
      return "Password should be at least 6 characters"
    }

    /**
     * This happened because on firebase configuration is not
     * allowed the specified provider example: email, fb, google...
     */
    if (error.code === "auth/operation-not-allowed") {
      return "The sing-in provider is disabled"
    }

    if (error.code === "auth/email-already-in-use") {
      return "The email address is already in use by another account"
    }

    if (error.code == "auth/wrong-password") {
      return "The password is invalid or the user does not have a password"
    }

    if (error.code == "auth/user-not-found") {
      return "here is no user record corresponding to this identifier. " +
        "The user may have been deleted"
    }

    return error.message
  }
}


