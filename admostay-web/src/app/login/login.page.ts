import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import {AuthenticationService} from "../core/services/authentication-service";
import { FirebaseErrorService } from "../core/services/firebase-errors"

@Component({
  selector: 'app-login',
  templateUrl: './login.page.html',
  styleUrls: ['./login.page.scss'],
})
export class LoginPage implements OnInit {

  credentials = {
    email: "",
    password: ""
  }

  showErrors : boolean = false;
  error: string = "";

  constructor(
    public authService: AuthenticationService,
    public firebaseErrorService: FirebaseErrorService,
    public router: Router
  ) { }

  ngOnInit() {}

  logIn() {
    this.authService.SignIn(this.credentials.email, this.credentials.password)
      .then( (res) => {
        console.error(res)
        this.router.navigate(['dashboard']);
      }).catch((error) => {
        console.log(error)
        this.error = this.firebaseErrorService.GetAuthError(error)
        this.showErrors = true
      })
  }
}
