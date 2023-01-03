import { Component } from '@angular/core';
import { ModalController } from '@ionic/angular';
import {PropertiesService} from "../core/services/properties-service";
import {LocalUserService} from "../core/services/user-local";
import {Property} from "../core/models/property";
import {HttpErrorResponse} from "@angular/common/http";
import { AlertController } from '@ionic/angular';
import {Router} from "@angular/router";

@Component({
  selector: 'app-create-property',
  templateUrl: './create-property.page.html',
  styleUrls: ['./create-property.page.scss'],
})
export class CreatePropertyPage  {

  /**
   * Property model object that will contain the information to
   * request.
   * */
  property:Property = {
    user_uid:  "",
    name: "",
    description: "",
    address: "",
    typ: "",
    rooms: 0,
  }

  /**
   * The show errors is a flag that will be in true
   * it there are errors on the creation process.
   * */
  showErrors: boolean = false;

  /**
   * This is the error message to show if there are errors.
   * */
  error: string = "";

  constructor(
    public router: Router,
    private modalController: ModalController,
    private alertController: AlertController,
    private localUser: LocalUserService,
    private properties:PropertiesService) { }

  /**
   * Cancel closes the modal component and cancel the attempt to
   * create a new property.
   * */
  cancel() {
    return this.modalController.dismiss(null, 'cancel');
  }

  /**
   * It sends to the backend the property object and store it.
   * */
  create() {
    this.property.user_uid = this.localUser.getCurrentUserUID()
    this.properties.createProperty(this.property)
      .subscribe( {
        next: (property) => this.handleResponse(property),
        error: (e) => this.handleErrors(e)
      })
  }

  async handleResponse(p:Property) {
    const alert = await this.alertController.create({
      header: 'Creation success',
      message: 'The property has been created',
      buttons: ['OK'],
    });
    await alert.present();
    await this.modalController.dismiss(null, 'ok');
    await this.router.navigateByUrl("/properties");

  }

  async handleErrors(e:HttpErrorResponse) {
    const message = `<div class="center validation-errors">
       <span>${e.error.message}</span>
    </div>`
    const alert = await this.alertController.create({
      message: message,
      buttons: ['OK'],
    });
    await alert.present();
  }
}
