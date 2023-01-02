import { Component } from '@angular/core';
import { ModalController } from '@ionic/angular';

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
  property = {
    name: "",
    description: "",
    address: "",
    type: "",
    rooms: 0,
  }

  constructor(private modalController: ModalController) { }

  /**
   * Cancel closes the modal component.
   * */
  cancel() {
    return this.modalController.dismiss(null, 'cancel');
  }

  /**
   * It sends to the backend the property object and store it.
   * */
  create() {

  }
}
