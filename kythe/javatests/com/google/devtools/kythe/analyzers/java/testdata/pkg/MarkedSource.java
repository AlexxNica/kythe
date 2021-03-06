//- @pkg ref Package
//- Package code PackageId
//- PackageId.kind "IDENTIFIER"
//- PackageId.pre_text "pkg"
package pkg;

//- @MarkedSource defines/binding Class
//- Class childof Package
//- Class code ClassId
//- ClassId child.0 ClassCxt
//- ClassCxt child.0 ClassCxtPackage
//- ClassCxtPackage.pre_text "pkg"
//- ClassId child.1 ClassTok
//- ClassTok.kind "IDENTIFIER"
//- ClassTok.pre_text "MarkedSource"
public class MarkedSource {

  //- @fieldName defines/binding Field
  //- Field childof Class
  //- @int ref Int
  //- Field code FieldId
  //- FieldId child.0 FieldCxt
  //- FieldCxt child.1 FieldCxtClass
  //- FieldCxtClass.pre_text "MarkedSource"
  //- FieldId child.1 FieldTok
  //- FieldTok.pre_text "fieldName"
  int fieldName;

  //- @methodName defines/binding Method
  //- Method childof Class
  //- Method typed MType
  //- MType param.1 Void
  //- Method code MethodId
  //- MethodId child.0 MethodCxt
  //- MethodCxt child.1 MethodCxtClass
  //- MethodCxtClass.pre_text "MarkedSource"
  //- MethodId child.1 MethodTok
  //- MethodTok.pre_text "methodName"
  //- MethodId child.2 MethodParams
  //- MethodParams.kind "PARAMETER_LOOKUP_BY_PARAM"
  //- MethodParams.pre_text "("
  //- MethodParams.post_child_text ", "
  //- MethodParams.post_text ")"
  void methodName() {}

  //- @methodWithParams defines/binding MethodWithParams
  //- MethodWithParams typed MPType
  //- MPType param.1 Void
  //- MPType param.2 String
  //- MPType param.3 Int
  void methodWithParams(String a, int b) {}

  //- @Inner defines/binding InnerClass
  //- InnerClass code InnerId
  public class Inner {
    //- @field defines/binding IField
    //- IField code CIField
    //- CIField child.0 CIFieldCxt
    //- CIFieldCxt child.0 CIFieldPkg
    //- CIFieldPkg.pre_text "pkg"
    //- CIFieldCxt child.1 CIFieldMS
    //- CIFieldMS.pre_text "MarkedSource"
    //- CIFieldCxt child.2 CIFieldInner
    //- CIFieldInner.pre_text "Inner"
    int field;
  }

  Object o = new Object() {
    //- @field defines/binding AnonField
    //- AnonField childof AnonClass
    //- AnonClass code AnonId
    //- AnonId child.0 AnonCxt
    //- AnonCxt.kind "CONTEXT"
    //- AnonCxt child.0 PkgToken
    //- PkgToken.pre_text pkg
    //- AnonCxt child.1 MksToken
    //- MksToken.pre_text "MarkedSource"
    //- AnonId child.1 AnonToken
    //- AnonToken.kind "IDENTIFIER"
    //- AnonToken.pre_text "(anon 1)"
    int field;
  };

  void methodWithInnerAnon() {
    Object o = new Object() {
      //- @field defines/binding IAField
      //- IAField code CIAField
      //- CIAField child.0 CIACxt
      //- CIACxt child.0 CIAPkg
      //- CIACxt child.1 CIAMs
      //- CIACxt child.2 CIAMethod
      //- CIACxt child.3 CIAAnon
      //- !{CIACxt child.4 _}
      //- CIAPkg.pre_text "pkg"
      //- CIAMs.pre_text "MarkedSource"
      //- CIAMethod.pre_text "methodWithInnerAnon"
      //- CIAAnon.pre_text "(anon 2)"
      int field;
    };
  }

  public class InnerAnon {
    Object o = new Object() {
      //- @field defines/binding IIAField
      //- IIAField code CIIAField
      //- CIIAField child.0 CIIACxt
      //- CIIACxt child.2 CIIAInnerCls
      //- CIIACxt child.3 CIIAInnerAnon
      //- CIIAInnerCls.pre_text "InnerAnon"
      //- CIIAInnerAnon.pre_text "(anon 1)"
      int field;
    };
  }

  //- Void code VoidId
  //- VoidId.kind "IDENTIFIER"
  //- VoidId.pre_text void
  //- Int code IntId
  //- IntId.kind "IDENTIFIER"
  //- IntId.pre_text int
  //- String code StringId
  //- StringId.kind "BOX"
  //- StringId child.0 StringCxt
  //- StringCxt.post_child_text "."
  //- StringCxt.add_final_list_token true
  //- StringId child.1 StringIdToken
  //- StringCxt.kind "CONTEXT"
  //- StringCxt child.0 JavaId
  //- JavaId.kind "IDENTIFIER"
  //- JavaId.pre_text java
  //- StringCxt child.1 LangId
  //- LangId.kind "IDENTIFIER"
  //- LangId.pre_text lang
  //- StringIdToken.kind "IDENTIFIER"
  //- StringIdToken.pre_text "String"
}
