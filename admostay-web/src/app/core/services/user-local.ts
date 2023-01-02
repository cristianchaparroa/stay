import {Injectable} from "@angular/core";
import {User} from "../models/user";

@Injectable({
  providedIn: 'root',
})
export class LocalUserService {

  constructor() {}

  getCurrentUser(): User | null {
    const userStored = localStorage.getItem('user')
    if (userStored) {
      const user = JSON.parse(userStored)
      return user;
    }
    return null;
  }

  getCurrentUserUID(): string {
    const user = this.getCurrentUser()
    let uid = ""
    if (user?.uid) {
      uid = user?.uid
    }
    return uid
  }
}
