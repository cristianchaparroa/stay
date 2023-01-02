import { Component, OnInit } from '@angular/core';
import {PropertiesService} from "../core/services/properties-service";
import {LocalUserService} from "../core/services/user-local";
import { ModalController } from '@ionic/angular';
import {CreatePropertyPage} from "../create-property/create-property.page";

@Component({
  selector: 'app-properties',
  templateUrl: './properties.page.html',
  styleUrls: ['./properties.page.scss'],
})
export class PropertiesPage implements OnInit {

  /**
   * List of properties that has a user.
   * */
  properties: any = []



  /**
   * @param propertiesService
   * @param localUser
   * */
  constructor(private propertiesService:PropertiesService,
              private localUser:LocalUserService,
              private modalController: ModalController) { }

  ngOnInit() {
    this.getAllProperties()
  }

  /**
   * It retrieves all properties and fill the properties array to be showed.
   * */
  getAllProperties(): void {
    const uid = this.localUser.getCurrentUserUID()
    this.propertiesService.getAllProperties(uid).subscribe(properties => {
      this.properties = properties
    })
  }

  /**
   * It handles how to show the create property modal
   * */
  async showCreateModal() {
    const modal = await this.modalController.create({
        component: CreatePropertyPage,
    });

    await modal.present();
  }

}
