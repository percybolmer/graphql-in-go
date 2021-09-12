
# Query with Filter arguemnts and Aliases
query {

  percentSoil: soil (percent:"2-3%") {
    name
  }
  soil (name:"Mulfattig") {
    name
    percent
  }
}



# Changes -- job folder
- gopher repository changes to the constructor and available job field

ONLY show getJobs since we havn't introduced ARGS