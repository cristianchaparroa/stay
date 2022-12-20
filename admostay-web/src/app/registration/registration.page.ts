import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import { AuthenticationService } from "../services/authentication-service";
import { FirebaseErrorService } from "../services/firebase-errors"

@Component({
  selector: 'app-registration',
  templateUrl: './registration.page.html',
  styleUrls: ['./registration.page.scss'],
})
export class RegistrationPage implements OnInit {

   credentials = {
    email: "",
    password: "",
  }
  showErrors : boolean = false;
  error: string = "";

  constructor(
    public authService: AuthenticationService,
    public firebaseErrorService: FirebaseErrorService,
    public router: Router) { }

  ngOnInit() {}

  signUp() {
    this.authService.RegisterUser(
      this.credentials.email,
      this.credentials.password
    )
      .then((res) => {
        this.showErrors = false;
        this.authService.SendVerificationMail()
        this.router.navigate(['verify-email']);
      }).catch((error) => {
        console.log(error.message)
        this.error = this.firebaseErrorService.GetAuthError(error);
        this.showErrors = true;
    })
  }
}
