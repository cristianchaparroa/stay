import { Component, OnInit } from '@angular/core';
import {PropertiesService} from '../core/services/properties-service'
import {LocalUserService} from "../core/services/user-local";
import {Router} from "@angular/router";

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.page.html',
  styleUrls: ['./dashboard.page.scss'],
})
export class DashboardPage implements OnInit {

  properties: any = []
  constructor( public router: Router,
               private propertiesService:PropertiesService,
              private localUser:LocalUserService) { }

  ngOnInit() {
    this.getAllProperties()
  }

  getAllProperties(): void {
    const  uid = this.localUser.getCurrentUserUID()
    this.propertiesService.getAllProperties(uid).subscribe(properties => {
      this.properties = properties
    })
  }
}
