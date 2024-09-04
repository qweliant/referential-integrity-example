Blurting method Pass 0

Ok so to me, this means, i take something and change it here, but the value doesnt update there. This may be bc we can update a value locally but not glabally bc of relationships elsewhere that can be affected by the graph being imcomplete. I willread a bit on this then come back

Referential integrity is the implied constraint placed on the tuples of values in a table referencing or pointing to values from another table
- ```
  Referential Integrity from R.A to S.B: Each Value in column A of table R must appear in the reference to column B of table S
  ```
- Example at 2:57 here
- Essentially, the foreign keys or multiple attribute foreign key relating tables must be referentially equal to each other, otherwise referential integrity cannot be validated
  - In PubPubs case, if we edit the fields of a Pub in one place they should be reflected across references to them.  
    - This is important bc of the proliferation of fields throughout our current data model. Not particularly the fields themselves but the values. 
    - If the value of a child is updated in one place, is that reflected across instantiations of the relationship?
      -  Before I can answer that i have to think about the examples of children and siblings in a pub
        - In the examples here we see children can be versions, images, Footnotes, etc. We also have siblings; tags, Collections, Contributors, etc. 
        - Some of these feel like pubTypes tho. A relationship is not a link to a download, nor is it a pubs version. Those seem like fields in a pubType. 
          - Which is why we should clarify how we think about relationships
        - Should we be saying a child or sibling is any collection of fields that can be shared between pubs? If so that seems too casual approach.
        - Like a tag can be applied to a pub outside of its fields to form a collection for instance. Or is a collection one Pub and its children the items in the Pub. I can see arguments for both fields and relationships and a relationships table. 
          - Like having a tag field opts you in and allows users to define their own Collections of media so either way works(disregard db constraints)
          - we can also experiment with users opt in for globally defined tags as mentioned by Gabe Stein way long ago. 
      - I like the idea of editing fields anywhere by reference. It would be cool if a field was global and could be updated across communities
        - like say the avg growth rate of squash become a value determined by review processes across communities type shit
    - I wonder if we can CAN enforce referential integrity through Json values. Maybe thats what Eric means when he says adding IDs to the blob of field values to act as a unique_id restraint?? Idk. Regardless it seems problematic because we would have to keep references in the db updated each time a relationship has something done on it and may not be able to do that easily on a jsonBlob. Or the values just could become a table maybe???? Regardless we need to flesh out relationships to understand how we should model them and what may or may not be possible. Seem slike a lot of uncertainty

There are a few ways to ruin referential integrity. If we think about the image above, we see APPLY references a unique foreign key to STUDENT and a unique column in COLLEGE
- Remember, for referential integrity from R.A referencing S.B each value in column A of table R must appear in column B of table S
  - A is the foreign key restraint. cName or sID in APPLY
  - B is the primary key for table S or least unique. cName in COLLEGE or sId in STUDENT 
  - Multiple attribute foreign keys are treated as A
    - where A would be a unique combination of fields in COLLEGE e.g cName_cState 
- Potential violations of referential integrity include
  - Insert into R
    - Insertion into R will violate the referential integrity bc neither exist in reference relation S
      - i.e. COLLEGE and STUDENT
  - Delete from S
    - Since A in R is a reference to B in S we cannot delete B from S without changing the schema of R
  - Update R.A
    - Updating the value of A shouldn't work either. These are references to values in S not the values themselves
  - Update S.B
    - We cannot update S.B bc that value is referenced as a foreign key in R.A. That one doesn't seem as obvious
- We can start to see how this is really a pointer thing. But I think theres relational db architecture i'm not aware of like CASCADING deletes etc.
  
Well, there are some special actions lol
- Delete from S(the reference table)
  - Restrict(default)
    - Do not allow the deletion
  - Set Null
    - If we delete a tuple in S we set the reference value in R to null. Whyyyyyyyyyyy. I mean I guess
  - Cascade
    - Deletes any tuple that references that value 
- Updates S.B(values in reference table)
  - Restrict(default)
    - Do not do the update
  - Set Null
    - Sets R.A to null
  - Cascade
    - propagates change to any referencing values.
    - if S.B changes any tuple referencing B will update its value
      - are cascading changes expensive?
Do these actions work in reverse? Can we update R.A and that change S.B
