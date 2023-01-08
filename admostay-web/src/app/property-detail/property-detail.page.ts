import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';
import {Property} from "../core/models/property";

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
    private modalController: ModalController) {
  }

  cancel() {
    return this.modalController.dismiss(null, 'cancel');
  }

}
