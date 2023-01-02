import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { CreatePropertyPage } from './create-property.page';

const routes: Routes = [
  {
    path: '',
    component: CreatePropertyPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class CreatePropertyPageRoutingModule {}
