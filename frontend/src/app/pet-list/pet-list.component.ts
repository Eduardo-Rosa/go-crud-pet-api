import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpClientModule } from '@angular/common/http'; // Importação necessária
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

const API_URL = 'http://go_pet_api:8080'; // URL da API


interface Pet {
  id: number;
  name: string;
  species: string;
  breed: string;
  age: number;
  birthDate: string;
  ownerName: string;
}

@Component({
  selector: 'app-pet-list',
  templateUrl: './pet-list.component.html',
  styleUrls: ['./pet-list.component.css'],
  standalone: true,
  imports: [CommonModule, FormsModule, HttpClientModule] // Adicionado HttpClientModule
})
export class PetListComponent implements OnInit {
  pets: Pet[] = [];
  pet: Pet = { id: 0, name: '', species: '', breed: '', age: 0, birthDate: '', ownerName: '' };
  isEditing = false;

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.loadPets();
  }

  editPet(pet: any): void {
    this.pet = { ...pet }; // Preenche o formulário com os dados do pet selecionado
    this.isEditing = true; // Define o modo de edição como verdadeiro
  }

  loadPets(): void {
    this.http.get<Pet[]>(`${API_URL}/pets`).subscribe({
      next: (data: Pet[]) => (this.pets = data),
      error: (error: any) => console.error('Erro ao carregar pets:', error)
    });
  }

  onSubmit(): void {
    if (this.isEditing) {
      this.http.put(`${API_URL}/pets/${this.pet.id}`, this.pet).subscribe({
        next: () => {
          this.loadPets();
          this.resetForm();
        },
        error: (error: any) => console.error('Erro ao atualizar pet:', error)
      });
    } else {
      this.http.post(`${API_URL}/pets`, this.pet).subscribe({
        next: () => {
          this.loadPets();
          this.resetForm();
        },
        error: (error: any) => console.error('Erro ao adicionar pet:', error)
      });
    }
  }

  deletePet(id: number): void {
    this.http.delete(`${API_URL}/pets/${id}`).subscribe({
      next: () => this.loadPets(),
      error: (error: any) => console.error('Erro ao deletar pet:', error)
    });
  }

  resetForm(): void {
    this.pet = { id: 0, name: '', species: '', breed: '', age: 0, birthDate: '', ownerName: '' };
    this.isEditing = false;
  }
}
