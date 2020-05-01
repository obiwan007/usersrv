export class TeamEntry {
  id: string = "0";
  name: string = "Default Team Name";
  member: MemberEntry[] = [];
  notes: string = "";
  constructor() {}
}

export class MemberEntry {
  id: string = "0";
  name: string = "Default Team Name";
  role: string;
  constructor() {}
}

export class Team {
  entries: TeamEntry[] = [];

  Add(p: TeamEntry) {
    this.entries.push(p);
  }

  Entries(): TeamEntry[] {
    return this.entries;
  }
}

export default new Team();
