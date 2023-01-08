import {Injectable} from "@angular/core";
import {HttpClient} from '@angular/common/http';
import { environment } from 'src/environments/environment';
import {Property} from "../models/property";
import {BaseService} from "./base-service";

@Injectable({
  providedIn: 'root',
})
export class PropertiesService extends BaseService {

  constructor(private http: HttpClient) {
    super();
  }

  /**
   * It will retrieve all user properties
   *
   * @param uid user identifier
   * */
  getAllProperties(uid:string) {
    const options = {headers: this.getBaseHeaders()};
    return this.http.get<Property[]>(`${environment.apiURL}/users/${uid}/properties`, options);
  }

  /**
   * It will add a new user property
   *
   * @param property body request to create
   * */
  createProperty(property:Property) {
    const options = {headers: this.getBaseHeaders()};
    return this.http.post<Property>(`${environment.apiURL}/users/properties`, property, options);
  }

  /**
   * It will delete a user property
   *
   * @param uid user identification
   * @param id property to delete
   * */
  deleteProperty(uid:string, id:string) {
    const options = {headers: this.getBaseHeaders()};
    return this.http.delete<Property[]>(`${environment.apiURL}/users/${uid}/properties/${id}`, options);
  }
}
