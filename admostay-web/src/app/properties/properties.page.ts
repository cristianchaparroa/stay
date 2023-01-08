import { Component, OnInit } from '@angular/core';
import {PropertiesService} from "../core/services/properties-service";
import {LocalUserService} from "../core/services/user-local";
import { ModalController } from '@ionic/angular';
import {CreatePropertyPage} from "../create-property/create-property.page";
import {Property} from "../core/models/property";
import {PropertyDetailPage} from "../property-detail/property-detail.page";

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
              private modalController: ModalController) {
    this.ngOnInit()
  }

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
    const { role } = await modal.onWillDismiss();
    if (role === "ok") {
      this.getAllProperties();
    }
  }

  async showDetail(p:Property) {
    const modal = await this.modalController.create({
      component:PropertyDetailPage,
      componentProps: {
        property:p,
      }
    })

    await  modal.present();
  }



}
