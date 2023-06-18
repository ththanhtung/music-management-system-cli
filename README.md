
# Music Management System - CLI
## Description

 The system allow users to organize and manage their music collection (tracks and playlists) by providing various features to add, update, search, and delete entries.

## ERD Model
![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/73ff1809-3ba1-42f5-8fb7-597cc73a07ae)

## Database Model
![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/3c0eb39e-b05b-4fcb-b947-8ebbbfa3ffaf)

## Note

In order to run this program you have to have setup and install Go Programming language in your machine
## Run Locally

If you have Go already setup and installed in your machine:
```bash
  go run main.go
```
If you have Docker already installed on your machine:
```bash
  docker pull thanhtung3172001/music-management-system-cli
```
```bash
  docker container run -it thanhtung3172001/music-management-system-cli
```
If you don't have Go already setup and install in your machine:
- Go to this website to install Go runtime: https://go.dev/dl/
- Then just run the program and press next
- Now you can run this command: ```go run main.go ```

## How to use
After enter the command you should see something like this:
![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/934f8d17-abf0-4afd-af0d-b87816298f23)
#### Add new music track:
- Press 1 
- Then add all the information needed for the track
- Example: 
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/a53d6fc9-b97f-487a-bb20-5d8313cc3f56)
- After fill out all the needed information press enter then the system will display the track info you have just added: 
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/3afe4737-ed72-4543-b94b-c09d33624468)
####  View details of a music track:
- press 2
- Enter the ID of the track that you want to view information:
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/e17e94f2-14db-4c52-9e85-80d7cb0ee646)
- Then the system will display all the information of the track:
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/e8e0193a-b9b5-4a5e-9583-cf0fde15a295)
#### Update details of a music track
- press 3
- Enter track's ID
- Enter the information of the track you want to update:
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/323a1185-1411-4a9f-ba77-cb09b39632fa)
- if the track is successfully update, the system will notify you:
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/df182c65-d02a-4aec-9bef-e74a81f72c03)
#### Delete a music track:
- press 4
- Enter track's ID
- If the track successfully deleted, the system will notify you: 
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/53b5ec3e-0ef4-4f62-a70d-12bea4fbd2dc)
####  Display all music tracks
- press 5
- The system then display all the song avaialble:
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/ff804007-d53e-4f9d-ab09-9f4b4171eef9)
 #### Create a playlist:
 - Press 6
 - Enter the playlist name:
 - ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/56a2bc43-a793-4133-9256-c287a363f1ec)
- if you successfully create a playlist, the system will notify you:
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/5b52d65b-1acb-488c-a987-3d87e4ec0303)
####  Add a track to a playlist:
- Press 7
- The system then display all the avaialble track you can add to new playlist, then just type which track you want to add a playlist:
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/e3e7a4f8-1c3f-4f76-8412-00f42a2617fe)
- then choose the playlist you want to add it to: 
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/fe36970d-783f-44fa-8e6f-4e2e9170e61e)
- If the track successfully added to the playlist, the system will notify you:
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/4fce39ab-a852-4574-9028-ab4c5219d206)

####  Search music tracks and playlists:
- Press 8
- Choose what what you what to search the track by (for example, I will search by title):
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/aac3fa8b-eabc-4269-874c-e69b18f9e0d6)
- Then enter the track's name: 
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/ca0e3e0e-fe1e-4a54-8005-96fb76313112)
- If the track is avaialble in the system's database, it will display the track name and which playlist the track is in(if it have): 
- ![image](https://github.com/ththanhtung/music-management-system-cli/assets/83943695/38ad2760-9328-4343-a7b7-109a5d63c404)