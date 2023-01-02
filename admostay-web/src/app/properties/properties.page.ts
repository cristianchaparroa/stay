import { Component, OnInit } from '@angular/core';
import {PropertiesService} from "../core/services/properties-service";
import {LocalUserService} from "../core/services/user-local";

@Component({
  selector: 'app-properties',
  templateUrl: './properties.page.html',
  styleUrls: ['./properties.page.scss'],
})
export class PropertiesPage implements OnInit {
  properties: any = []
  constructor(private propertiesService:PropertiesService,
              private localUser:LocalUserService) { }

  ngOnInit() {
    this.getAllProperties()
  }

  getAllProperties(): void {
    const uid = this.localUser.getCurrentUserUID()
    this.propertiesService.getAllProperties(uid).subscribe(properties => {
      this.properties = properties
    })
  }

}
