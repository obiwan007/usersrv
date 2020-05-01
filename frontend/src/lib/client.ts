let index = 0;
export class ClientEntry {
  id: string = "0";
  team: string = "";
  status: string = "";
  name: string = "Default Client";
  adress: string = "";
  elapsedSeconds: number = 0;
  notes: string = "";
  constructor() {}

  static fromRaw(
    entry: any,
    existingEntry: ClientEntry = new ClientEntry()
  ): ClientEntry {
    return Object.assign(existingEntry, entry);
  }
}

export class Client {
  constructor() {
    const d = localStorage.getItem("clients");
    if (d) {
      this.entries =
        d && JSON.parse(d).map((entry: any) => ClientEntry.fromRaw(entry));
    }
    const i = localStorage.getItem("clientsIndex");

    index = i ? +i : 0;
  }
  save() {
    localStorage.setItem("clients", JSON.stringify(this.entries));
    localStorage.setItem("clientsIndex", JSON.stringify(index));
  }
  entries: ClientEntry[] = [];

  add(p: ClientEntry) {
    p.id = "" + index++;
    this.entries.push(ClientEntry.fromRaw(p));
    this.save();
    return this.Entries();
  }

  del(p: ClientEntry) {
    let data = this.entries;
    const entry = data.find((d) => d.id === p.id);

    if (entry) {
      console.log("Delete ", p);
      const index = data.indexOf(entry);
      data.splice(index, 1);
    }
    this.save();
    return this.Entries();
  }

  update(p: ClientEntry) {
    let data = this.entries;
    const entry = data.find((d) => d.id === p.id);

    if (entry) {
      console.log("Delete ", p);
      ClientEntry.fromRaw(p, entry);
    }
    this.save();
    return this.Entries();
  }

  Entries(): ClientEntry[] {
    return [...this.entries];
  }

  EntriesDict(): any {
    const dict: any = {};
    this.entries.forEach((e) => (dict[+e.id] = e.name));
    return dict;
  }
}

export default new Client();
