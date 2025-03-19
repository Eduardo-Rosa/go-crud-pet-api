import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { Routes } from '@angular/router';
import { PetListComponent } from './pet-list/pet-list.component';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

export const routes: Routes = [
  { path: '', component: PetListComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes),
    CommonModule,
    FormsModule,
  ],
  exports: [RouterModule],
})
export class AppRoutingModule { }