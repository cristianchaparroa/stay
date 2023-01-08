import {Injectable} from "@angular/core";

@Injectable({
  providedIn: 'root',
})
export class FirebaseErrorService {

   constructor() {}

   public GetAuthError(error: { code: string; message: string }) {
    const pattern = "\\(auth.*";
    const regex = new RegExp( pattern);
    if (regex.test(error.message)) {
      error.message = error.message.replace(regex, "");
    }
    error.message = error.message.replace("Firebase:", "");
    return error.message;
  }
}


