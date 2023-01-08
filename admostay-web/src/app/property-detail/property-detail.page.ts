import { Component, OnInit } from '@angular/core';
import {AlertController, ModalController} from '@ionic/angular';
import {Property} from "../core/models/property";
import {PropertiesService} from "../core/services/properties-service";
import {HttpErrorResponse} from "@angular/common/http";
import {Router} from "@angular/router";

@Component({
  selector: 'app-property-detail',
  templateUrl: './property-detail.page.html',
  styleUrls: ['./property-detail.page.scss'],
})
export class PropertyDetailPage {
  /**
   * Property to show the details.
   * */
  property:Property = {
    id: "",
    user_uid: "",
    name: "",
    description: "",
    address: "",
    typ: "",
    rooms: 0,
  }

  constructor(
    public router: Router,
    private modalController: ModalController,
    private propertiesService:PropertiesService,
    private alertController: AlertController,
  ) {
  }

  cancel() {
    return this.modalController.dismiss(null, 'cancel');
  }

  async delete() {
    const alert = await this.alertController.create({
      message: 'Are you sure you want to delete the property?',
      buttons: [
        {
          text: 'Cancel',
          role: 'cancel',
        },
        {
          text: 'continue',
          role: 'continue',
        },
      ],
    });

    await alert.present();
    const { role } = await alert.onDidDismiss();
    if (role == 'cancel') {
      return
    }

    this.propertiesService
      .deleteProperty(this.property.user_uid, this.property.id)
      .subscribe({
        next: () => this.handleDeleteResponse(),
        error: (e) => this.handleDeleteErrors(e),
      })
  }

  async handleDeleteResponse() {
    const alert = await this.alertController.create({
      header: 'Delete success',
      message: 'The property has been deleted',
      buttons: ['OK'],
    });
    await alert.present();
    await this.modalController.dismiss(null, 'ok');
    await this.router.navigateByUrl("/properties");
  }

  async handleDeleteErrors(e:HttpErrorResponse) {
    console.error(e);
  }

}
