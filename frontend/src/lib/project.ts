let index = 0;
export class ProjectEntry {
  id: string = "0";
  team: string = "";
  status: string = "";
  client: string = "";
  project: string = "Default Project";
  elapsedSeconds: number = 0;
  description: string = "";
  constructor() {}
  static fromRaw(
    entry: any,
    existingEntry: ProjectEntry = new ProjectEntry()
  ): ProjectEntry {
    return Object.assign(existingEntry, entry);
  }
}

export class Project {
  entries: ProjectEntry[] = [];
  /**
   *
   */
  constructor() {
    const d = localStorage.getItem("projects");
    if (d) {
      this.entries =
        d && JSON.parse(d).map((entry: any) => ProjectEntry.fromRaw(entry));
    }
    const i = localStorage.getItem("projectsIndex");

    index = i ? +i : 0;
  }
  save() {
    localStorage.setItem("projects", JSON.stringify(this.entries));
    localStorage.setItem("projectsIndex", JSON.stringify(index));
  }

  addProject(p: ProjectEntry) {
    this.entries.push(p);
    this.save();
  }
  add(p: ProjectEntry) {
    p.id = "" + index++;
    this.entries.push(ProjectEntry.fromRaw(p));
    this.save();
    return this.Entries();
  }

  del(p: ProjectEntry) {
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

  update(p: ProjectEntry) {
    let data = this.entries;
    const entry = data.find((d) => d.id === p.id);

    if (entry) {
      console.log("Delete ", p);
      ProjectEntry.fromRaw(p, entry);
    }
    this.save();
    return this.Entries();
  }

  Entries(): ProjectEntry[] {
    return [...this.entries];
  }
}

export default new Project();
