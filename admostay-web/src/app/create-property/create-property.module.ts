import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { CreatePropertyPageRoutingModule } from './create-property-routing.module';

import { CreatePropertyPage } from './create-property.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    CreatePropertyPageRoutingModule
  ],
  declarations: [CreatePropertyPage]
})
export class CreatePropertyPageModule {}
