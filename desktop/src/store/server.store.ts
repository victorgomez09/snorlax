import { create } from "zustand";
import { ServerType } from "@/types/server.type";
import { deleteDestination, getAllDestinations } from "@/api/destination.api";

interface ServerStore {
  loading: boolean;
  servers: ServerType[];
  selectedServer: ServerType | null;
  setSelectedServer: (selectedServer: ServerType | null) => void;
  loadServers: (selectServerWithId?: number) => void;
  deleteServer: (serverId: number, serverName: string) => void;
  editServer: (serverId: number, connection: string, name: string) => void;
}

export const useServerStore = create<ServerStore>((set) => ({
  loading: false,
  servers: [],
  selectedServer: null,
  setSelectedServer: (selectedServer) => {
    // set selected database id to localstorage
    if (selectedServer) {
      localStorage.setItem(
        "selectedServerId",
        JSON.stringify(selectedServer.id)
      );
    }

    set({ selectedServer });
  },
  loadServers: (selectServerWithId) => {
    // get serverId from localstorage is `selectedServerId` is not given
    let servers: { id: number; name: string; url: string }[] = [];
    getAllDestinations().then((data) => {
      set({ servers: data.data || [] });
      servers = data.data;
    });
    if (!selectServerWithId) {
      const server = localStorage.getItem("selectedServerId");
      selectServerWithId = server ? JSON.parse(server) : {};
    }

    // select server if id is given
    if (selectServerWithId) {
      // check if database exist in loaded data
      let filteredServer = servers.filter(
        (server: ServerType) => server.id === selectServerWithId
      );

      if (filteredServer.length == 1) {
        set({ selectedServer: filteredServer[0] });
      }
    }

    // invoke('read_servers')
    //   .then((servers: any) => {
    //     set({ servers, loading: false });

    //     // get serverId from localstorage is `selectedServerId` is not given
    //     if (!selectServerWithId) {
    //       selectServerWithId = JSON.parse(
    //         localStorage.getItem('selectedServerId') || ''
    //       );
    //     }

    //     // select server if id is given
    //     if (selectServerWithId) {
    //       // check if database exist in loaded data
    //       let filteredServer = servers.filter(
    //         (server: ServerType) => server.id === selectServerWithId
    //       );

    //       if (filteredServer.length == 1) {
    //         set({ selectedServer: filteredServer[0] });
    //       }
    //     }
    //   })
    // .catch(console.log);
  },
  editServer(serverId, connection, name) {
    set((state) => ({
      servers: state.servers.map((server) => {
        if (serverId === server.id) {
          server.name = name;
          server.url = connection;
        }

        return server;
      }),
    }));
  },
  deleteServer(serverId, serverName) {
    deleteDestination(serverName).then(() => {
      set((state) => ({
        servers: state.servers.filter((server) => server.id !== serverId),
        selectedServer:
          state.selectedServer?.id == serverId
            ? state.servers.length > 1
              ? state.servers[0]
              : null
            : state.selectedServer,
      }));
    });
  },
}));
