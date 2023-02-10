export interface ListInitStudent{
    StudentId:number,
    ClassId:number,
    Year:number,
    StudentName:string,
    Gender:string,
    Birthday:string,
    // UserId:number
}
interface selectDataInt{
    stuId:number,
    page:number,
    count:number,
    pageSize:number,
}

export class InitStu{
    selectData:selectDataInt={
        stuId:201764712,
        page:1,
        count:0,
        pageSize:10
    }
    list:ListInitStudent[] = []
}