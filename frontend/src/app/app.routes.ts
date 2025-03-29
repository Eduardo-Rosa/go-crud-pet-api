import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { PetListComponent } from './pet-list/pet-list.component';


export const routes: Routes = [
  { path: '', component: PetListComponent },
];

@NgModule({
  imports: [FormsModule, CommonModule, RouterModule.forRoot(routes)],
  exports: [RouterModule],
  
})

export class AppRoutingModule { }
export class AppModule {}